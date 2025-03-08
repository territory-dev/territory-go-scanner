package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/scanner"
	"go/token"
	"go/types"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode"

	pb "github.com/pawelstiasny/territory/indexers/go/pb"
	"google.golang.org/protobuf/encoding/protodelim"
	"google.golang.org/protobuf/proto"
)

type Package struct {
	Dir           string // directory containing package sources
	ImportPath    string // import path of package in dir
	ImportComment string // path in import comment on package statement
	Name          string // package name
	Doc           string // package documentation string
	Target        string // install path
	Shlib         string // the shared library that contains this package (only set when -linkshared)
	Goroot        bool   // is this package in the Go root?
	Standard      bool   // is this package part of the standard Go library?
	Stale         bool   // would 'go install' do anything for this package?
	StaleReason   string // explanation for Stale==true
	Root          string // Go root or Go path dir containing this package
	ConflictDir   string // this directory shadows Dir in $GOPATH
	BinaryOnly    bool   // binary-only package (no longer supported)
	ForTest       string // package is only for use in named test
	Export        string // file containing export data (when using -export)
	BuildID       string // build ID of the compiled package (when using -export)
	//Module         *Module  // info about package's containing module, if any (can be nil)
	Match          []string // command-line patterns matching this package
	DepOnly        bool     // package is only a dependency, not explicitly listed
	DefaultGODEBUG string   // default GODEBUG setting, for main packages

	// Source files
	GoFiles           []string // .go source files (excluding CgoFiles, TestGoFiles, XTestGoFiles)
	CgoFiles          []string // .go source files that import "C"
	CompiledGoFiles   []string // .go files presented to compiler (when using -compiled)
	IgnoredGoFiles    []string // .go source files ignored due to build constraints
	IgnoredOtherFiles []string // non-.go source files ignored due to build constraints
	CFiles            []string // .c source files
	CXXFiles          []string // .cc, .cxx and .cpp source files
	MFiles            []string // .m source files
	HFiles            []string // .h, .hh, .hpp and .hxx source files
	FFiles            []string // .f, .F, .for and .f90 Fortran source files
	SFiles            []string // .s source files
	SwigFiles         []string // .swig files
	SwigCXXFiles      []string // .swigcxx files
	SysoFiles         []string // .syso object files to add to archive
	TestGoFiles       []string // _test.go files in package
	XTestGoFiles      []string // _test.go files outside package

	// Embedded files
	EmbedPatterns      []string // //go:embed patterns
	EmbedFiles         []string // files matched by EmbedPatterns
	TestEmbedPatterns  []string // //go:embed patterns in TestGoFiles
	TestEmbedFiles     []string // files matched by TestEmbedPatterns
	XTestEmbedPatterns []string // //go:embed patterns in XTestGoFiles
	XTestEmbedFiles    []string // files matched by XTestEmbedPatterns

	// Cgo directives
	CgoCFLAGS    []string // cgo: flags for C compiler
	CgoCPPFLAGS  []string // cgo: flags for C preprocessor
	CgoCXXFLAGS  []string // cgo: flags for C++ compiler
	CgoFFLAGS    []string // cgo: flags for Fortran compiler
	CgoLDFLAGS   []string // cgo: flags for linker
	CgoPkgConfig []string // cgo: pkg-config names

	// Dependency information
	Imports      []string          // import paths used by this package
	ImportMap    map[string]string // map from source import to ImportPath (identity entries omitted)
	Deps         []string          // all (recursively) imported dependencies
	TestImports  []string          // imports from TestGoFiles
	XTestImports []string          // imports from XTestGoFiles

	// Error information
	Incomplete bool // this package or a dependency has an error
}

type ModuleInfo struct {
	Path string
	Dir  string
}

// var chrootDir = flag.String("chroot", "", "chroot before scanning")
// var setGid = flag.Int("setgid", 0, "set group ID before scanning")
// var setUid = flag.Int("setuid", 0, "set user ID before scanning")
// var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var system = flag.Bool("system", false, "include non-local dependencies")

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		fmt.Println("Usage: program <input_repository> <output_dir>")
		os.Exit(1)
	}
	/*
		if *cpuprofile != "" {
			f, err := os.Create(*cpuprofile)
			if err != nil {
				log.Fatal(err)
			}
			pprof.StartCPUProfile(f)
			defer pprof.StopCPUProfile()
		}

		if *chrootDir != "" {
			err := unix.Chroot(*chrootDir)
			if err != nil {
				fmt.Printf("failed to chroot: %v\n", err)
				os.Exit(1)
			}
		}
		if *setGid != 0 {
			err := unix.Setgid(*setGid)
			if err != nil {
				fmt.Printf("failed to setgid: %v\n", err)
				os.Exit(1)
			}
		}
		if *setUid != 0 {
			err := unix.Setuid(*setUid)
			if err != nil {
				fmt.Printf("failed to setuid: %v\n", err)
				os.Exit(1)
			}
		}*/
	repoDir, err := filepath.Abs(args[0])
	if err != nil {
		log.Fatal("failed to resolve repo path:", err)
	}
	if !strings.HasSuffix(repoDir, "/") {
		repoDir += "/"
	}
	err = os.Chdir(repoDir)
	if err != nil {
		log.Fatal("could not chdir to repo:", err)
	}

	outputDir := args[1]
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	nodeFile := outputDir + "/nodes.uim"
	indexFile := outputDir + "/search.uim"

	// Find all packages in the repository
	packages, err := findPackages(repoDir)
	if err != nil {
		log.Fatalf("Failed to find packages: %v", err)
	}

	log.Printf("Found %d packages", len(packages))

	// Open output files
	nodesOut, err := os.Create(nodeFile)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer nodesOut.Close()

	indexOut, err := os.Create(indexFile)
	if err != nil {
		log.Fatalf("Failed to create index file: %v", err)
	}
	defer indexOut.Close()

	nodesWriter := bufio.NewWriter(nodesOut)
	indexWriter := bufio.NewWriter(indexOut)
	defer nodesWriter.Flush()
	defer indexWriter.Flush()

	err = processAndStreamPackages(repoDir, packages, nodesWriter, indexWriter)
	if err != nil {
		log.Fatalf("Error processing packages: %v", err)
	}
}

type importerFunc func(path string) (*types.Package, error)

func (f importerFunc) Import(path string) (*types.Package, error) { return f(path) }

func findPackages(repoDir string) ([]Package, error) {
	var packages []Package
	found := make(map[string]struct{})

	err := filepath.WalkDir(repoDir, func(p string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			return nil
		}

		if _, ye := found[p]; ye {
			return nil
		}

		gomodPath := filepath.Join(p, "go.mod")
		if _, err := os.Stat(gomodPath); errors.Is(err, os.ErrNotExist) {
			return nil
		}

		log.Printf("Collecting packages: %s\n", p)

		cmd := exec.Command("go", "list", "-buildvcs=false", "-deps", "-json", "all")
		cmd.Dir = p

		out, err := cmd.Output()
		if err != nil {
			fmt.Printf("%s: go list failed: %v\n", p, err)
			if ee, ye := err.(*exec.ExitError); ye {
				fmt.Println(string(ee.Stderr))
			}
			return nil
		}

		decoder := json.NewDecoder(bytes.NewReader(out))

		for decoder.More() {
			var pkg Package
			if err := decoder.Decode(&pkg); err != nil {
				log.Printf("Error decoding package info: %v\n", err)
				log.Printf("go list output: %s\n", out)
				return nil
			}

			if _, ye := found[pkg.Dir]; !ye {
				packages = append(packages, pkg)
				found[pkg.Dir] = struct{}{}
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return packages, nil
}

func processAndStreamPackages(repoDir string, packages []Package, nodesWriter, indexWriter io.Writer) error {
	// Create a single FileSet for all files
	os.Chdir(repoDir)
	fset := token.NewFileSet()

	// Group files by package path for type checking
	type pkginfo struct {
		files      []*ast.File
		test_files []*ast.File
		pkg        *Package
	}
	packageInfos := make(map[string]pkginfo)
	fileContents := make(map[string][]byte)

	// Parse all files in the module first
	log.Println("parsing")
	for _, pkg := range packages {
		allFiles := make([]string, 0)
		allFiles = append(allFiles, pkg.GoFiles...)
		allFiles = append(allFiles, pkg.CgoFiles...)
		allFiles = append(allFiles, pkg.TestGoFiles...)
		allTestFiles := make([]string, 0)
		allTestFiles = append(allTestFiles, pkg.XTestGoFiles...)
		pi := pkginfo{
			pkg:        &pkg,
			files:      make([]*ast.File, 0, len(allFiles)),
			test_files: make([]*ast.File, 0, len(allTestFiles)),
		}
		for _, grp := range []struct {
			isTest bool
			files  []string
		}{
			{isTest: false, files: allFiles},
			{isTest: true, files: allTestFiles},
		} {
			for _, goFile := range grp.files {
				fullPath := filepath.Join(pkg.Dir, goFile)

				src, err := ioutil.ReadFile(fullPath)
				if err != nil {
					log.Printf("Failed to read file %s: %v", fullPath, err)
					continue
				}

				f, err := parser.ParseFile(fset, fullPath, src, parser.ParseComments)
				if err != nil {
					log.Printf("Failed to parse file %s: %v", fullPath, err)
					continue
				}

				if grp.isTest {
					pi.test_files = append(pi.test_files, f)
				} else {
					pi.files = append(pi.files, f)
				}

				fileContents[fullPath] = src
			}
		}
		packageInfos[pkg.Dir] = pi
	}

	conf := types.Config{
		Importer: importer.ForCompiler(fset, "source", nil),
		Error: func(err error) {
			log.Printf("Type error: %v", err)
		},
	}

	// Type check and process each package in the module
	pkgCtr := 0
	for pkgDir, pi := range packageInfos {
		pkgCtr++
		inRepo := strings.HasPrefix(pkgDir+"/", repoDir)
		if !inRepo && !*system {
			log.Printf("skipping system package: %s", pkgDir)
			continue
		}
		log.Printf("package [%d/%d]: %s", pkgCtr, len(packages), pkgDir)
		os.Chdir(pkgDir)

		info := &types.Info{
			Types: make(map[ast.Expr]types.TypeAndValue),
			Defs:  make(map[*ast.Ident]types.Object),
			Uses:  make(map[*ast.Ident]types.Object),
		}

		// regular .go files
		_, err := conf.Check(pi.pkg.ImportPath, fset, pi.files, info)
		if err != nil {
			log.Printf("Type checking error in package %s: %v", pkgDir, err)
		}

		for _, f := range pi.files {
			filename := fset.File(f.Pos()).Name()
			if src, ok := fileContents[filename]; ok {
				err := processAndStreamFile(repoDir, fset, f, info, src, nodesWriter, indexWriter)
				if err != nil {
					log.Printf("Error processing file %s: %v", filename, err)
				}
			}
		}

		// test files
		_, err = conf.Check(pi.pkg.ImportPath+"_test", fset, pi.test_files, info)
		if err != nil {
			log.Printf("Type checking error in package %s: %v", pkgDir, err)
		}

		for _, f := range pi.test_files {
			filename := fset.File(f.Pos()).Name()
			if src, ok := fileContents[filename]; ok {
				err := processAndStreamFile(repoDir, fset, f, info, src, nodesWriter, indexWriter)
				if err != nil {
					log.Printf("Error processing file %s: %v", filename, err)
				}
			}
		}
	}

	return nil
}

func processAndStreamFile(repoDir string, fset *token.FileSet, f *ast.File, info *types.Info, src []byte, nodesWriter, indexWriter io.Writer) error {
	var fileNodeWriter TokenWriter

	for _, decl := range f.Decls {
		var node *pb.Node
		switch d := decl.(type) {
		case *ast.FuncDecl:
			node = generateNode(repoDir, fset, d, info, src)
			if node != nil {
				if err := writeNodeToStream(node, nodesWriter); err != nil {
					return fmt.Errorf("failed to write node: %v", err)
				}

				// Generate and write index item for function
				indexItem := generateIndexItem(d.Name.Name, node, pb.IndexItemKind_IISymbol)
				if err := writeIndexItemToStream(indexItem, indexWriter); err != nil {
					return fmt.Errorf("failed to write index item: %v", err)
				}

				appendFunctionToFileNode(&fileNodeWriter, node)
			}
		case *ast.GenDecl:
			// fmt.Printf("%#v\n", d)
			if d.Tok == token.VAR || d.Tok == token.CONST || d.Tok == token.TYPE {
				node = generateNode(repoDir, fset, d, info, src)
				if node != nil {
					if err := writeNodeToStream(node, nodesWriter); err != nil {
						return fmt.Errorf("failed to write node: %v", err)
					}

					// Generate and write index items for each declaration
					for _, spec := range d.Specs {
						var name string
						var kind pb.IndexItemKind
						switch s := spec.(type) {
						case *ast.TypeSpec:
							name = s.Name.Name
							kind = pb.IndexItemKind_IISymbol
						case *ast.ValueSpec:
							for _, ident := range s.Names {
								name = ident.Name
								kind = pb.IndexItemKind_IISymbol
							}
						}
						if name != "" {
							indexItem := generateIndexItem(name, node, kind)
							if err := writeIndexItemToStream(indexItem, indexWriter); err != nil {
								return fmt.Errorf("failed to write index item: %v", err)
							}
						}
					}

					appendDeclToFileNodeInfo(&fileNodeWriter, node, d)
				}
			}
		}
	}

	fileNodeTokens, fileNodeText := fileNodeWriter.Get()
	fileNode := generateFileNode(repoDir, fset.File(f.Pos()).Name(), fileNodeTokens, fileNodeText)
	err := writeNodeToStream(fileNode, nodesWriter)
	if err != nil {
		return fmt.Errorf("failed to write file node: %v", err)
	}
	return nil
}

func generateIndexItem(name string, node *pb.Node, kind pb.IndexItemKind) *pb.IndexItem {
	return &pb.IndexItem{
		Key:  name,
		Kind: kind,
		Path: &node.Path,
		Href: &pb.IndexItem_UniHref{
			UniHref: &pb.UniHref{
				Path:   node.Path,
				Offset: node.Start.Offset,
			},
		},
	}
}

func generateFileNode(repoDir, filename string, tokens []*pb.Token, text string) *pb.Node {
	return &pb.Node{
		Id:     uint64(0),
		Kind:   pb.NodeKind_SourceFile,
		Path:   strings.TrimPrefix(filename, repoDir),
		Start:  &pb.Location{Line: 0, Column: 0, Offset: 0},
		Text:   text,
		Tokens: tokens,
	}
}

func writeNodeToStream(node *pb.Node, writer io.Writer) error {
	data, err := proto.Marshal(node)
	if err != nil {
		return err
	}

	// Write the length of the marshaled data as a varint
	length := uint64(len(data))
	lengthBuf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(lengthBuf, length)
	_, err = writer.Write(lengthBuf[:n])
	if err != nil {
		return err
	}

	// Write the marshaled data
	_, err = writer.Write(data)
	return err
}

func writeIndexItemToStream(item *pb.IndexItem, writer io.Writer) error {
	_, err := protodelim.MarshalTo(writer, item)

	return err
}

func generateNode(repoDir string, fset *token.FileSet, decl ast.Node, info *types.Info, src []byte) *pb.Node {
	var kind pb.NodeKind
	var start, end token.Pos

	switch d := decl.(type) {
	case *ast.FuncDecl:
		kind = pb.NodeKind_Definition
		if d.Doc != nil {
			start = d.Doc.Pos()
		} else {
			start = d.Pos()
		}
		end = d.End()
	case *ast.GenDecl:
		if len(d.Specs) > 0 {
			switch d.Specs[0].(type) {
			case *ast.TypeSpec:
				kind = pb.NodeKind_Definition
			case *ast.ValueSpec:
				kind = pb.NodeKind_Definition
			default:
				return nil // Skip other kinds of declarations
			}
		} else {
			return nil // Skip empty declarations
		}
		if d.Doc != nil {
			start = d.Doc.Pos()
		} else {
			start = d.Pos()
		}
		end = d.End()
	default:
		return nil
	}

	startPos := fset.Position(start)
	endPos := fset.Position(end)

	// Extract the full text of the declaration
	nodeText := string(src[startPos.Offset:endPos.Offset])

	node := &pb.Node{
		Id:   uint64(start),
		Kind: kind,
		Path: strings.TrimPrefix(fset.File(start).Name(), repoDir),
		Start: &pb.Location{
			Line:   uint32(startPos.Line),
			Column: uint32(startPos.Column),
			Offset: uint32(startPos.Offset),
		},
		Text:                nodeText,
		UimReferenceContext: UimReferenceContext(decl),
	}

	// Generate tokens
	node.Tokens = generateTokens(repoDir, fset, decl, src, start, end, startPos.Offset, info)

	return node
}

func generateTokens(repoDir string, fset *token.FileSet, decl ast.Node, src []byte, start, end token.Pos, nodeStartOffset int, info *types.Info) []*pb.Token {
	var tokens []*pb.Token
	var s scanner.Scanner
	file := fset.File(start)
	s.Init(file, src, nil, scanner.ScanComments)

	// Create a map of positions to identifiers
	posToIdent := make(map[token.Pos]*ast.Ident)
	ast.Inspect(decl, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok {
			posToIdent[ident.Pos()] = ident
		}
		return true
	})

	lastOffset := fset.Position(start).Offset
	for {
		pos, tok, _ := s.Scan()
		if pos >= end {
			break
		}
		if pos < start {
			continue
		}

		position := fset.Position(pos)
		curOffset := position.Offset
		wsStart := position.Offset
		for wsStart > lastOffset {
			if !unicode.IsSpace(rune(src[wsStart-1])) {
				break
			}
			wsStart--
		}

		// Add whitespace token if there's a gap
		if wsStart > lastOffset && wsStart < curOffset {
			tokens = append(tokens, &pb.Token{
				Offset:      uint32(wsStart - nodeStartOffset),
				Type:        pb.TokenType_WS,
				UimLocation: nil,
			})
		}

		tokenType := mapTokenType(tok)
		token := &pb.Token{
			Offset: uint32(curOffset - nodeStartOffset),
			Type:   tokenType,
			UimLocation: &pb.Location{
				Line:   *proto.Uint32(uint32(position.Line)),
				Column: *proto.Uint32(uint32(position.Column)),
				Offset: *proto.Uint32(uint32(position.Offset)),
			},
		}

		if ident, ok := posToIdent[pos]; ok {
			token.Href = makeHref(info, fset, repoDir, ident)
		}

		tokens = append(tokens, token)
		lastOffset = curOffset
	}

	return tokens
}

func makeHref(info *types.Info, fset *token.FileSet, repoDir string, ident *ast.Ident) *pb.Token_UniHref {
	obj := info.ObjectOf(ident)
	if obj == nil {
		return nil
	}

	// ignore package names as those point to the import block
	if _, ye := obj.(*types.PkgName); ye {
		return nil
	}

	// embedded struct
	if v, ye := obj.(*types.Var); ye && v.Embedded() {
		// "For an embedded field, Uses returns the *TypeName it denotes"
		// <https://pkg.go.dev/go/types#Var.Embedded>
		obj = info.Uses[ident]
		if obj == nil {
			return nil
		}
	}

	defPos := fset.PositionFor(obj.Pos(), false)
	if defPos.Filename == "" {
		return nil
	}

	return &pb.Token_UniHref{UniHref: &pb.UniHref{
		Path:   strings.TrimPrefix(defPos.Filename, repoDir),
		Offset: uint32(defPos.Offset),
	}}
}

func mapTokenType(tok token.Token) pb.TokenType {
	switch {
	case tok.IsKeyword():
		return pb.TokenType_Keyword
	case tok.IsOperator():
		return pb.TokenType_Punctuation
	case tok == token.IDENT:
		return pb.TokenType_Identifier
	case tok == token.COMMENT:
		return pb.TokenType_Comment
	case tok.IsLiteral():
		return pb.TokenType_Literal
	default:
		return pb.TokenType_WS
	}
}
