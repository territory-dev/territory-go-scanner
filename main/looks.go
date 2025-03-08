package main

import (
	"go/ast"
	"go/token"
	"strings"

	"github.com/golang/protobuf/proto"
	pb "github.com/pawelstiasny/territory/indexers/go/pb"
)

func appendFunctionToFileNode(tw *TokenWriter, node *pb.Node) {
	var lineCtr int
	href := &pb.UniHref{
		Path:   node.Path,
		Offset: node.Start.Offset,
	}

	for i, token := range node.Tokens {
		startOffset := token.Offset
		endOffset := uint32(len(node.Text))
		if i < len(node.Tokens)-1 {
			endOffset = node.Tokens[i+1].Offset
		}

		tokenText := node.Text[startOffset:endOffset]

		tw.AppendWithReal(
			token.Type, tokenText,
			proto.Uint32(node.Start.Offset+token.Offset),
			proto.Uint32(node.Start.Line+uint32(lineCtr)),
			href)

		lineCtr += strings.Count(tokenText, "\n")

		if tokenText == "{" && token.Type == pb.TokenType_Punctuation {
			tw.Append(pb.TokenType_WS, " … ", href)
			tw.Append(pb.TokenType_Punctuation, "}", href)
			break
		}
	}
	tw.Append(pb.TokenType_WS, "\n\n", nil)
}

func tokStr(tok token.Token) string {
	switch tok {
	case token.CONST:
		return "const"
	case token.TYPE:
		return "type"
	case token.VAR:
		return "var"
	}
	return ""
}

func appendDeclToFileNodeInfo(tw *TokenWriter, node *pb.Node, decl *ast.GenDecl) {
	var lineCtr int
	href := &pb.UniHref{
		Path:   node.Path,
		Offset: node.Start.Offset,
	}

	kw := tokStr(decl.Tok)

	for _, spec := range decl.Specs {
		off := node.Start.Offset

		switch spec := spec.(type) {
		case *ast.ValueSpec:
			if len(spec.Names) == 1 {
				tw.AppendWithReal(pb.TokenType_Keyword, kw, proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
				tw.AppendWithReal(pb.TokenType_WS, " ", proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
				tw.AppendWithReal(pb.TokenType_Identifier, spec.Names[0].Name, proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
				tw.AppendWithReal(pb.TokenType_WS, " …", proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
			} else {
				tw.AppendWithReal(pb.TokenType_Keyword, kw, proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
				tw.AppendWithReal(pb.TokenType_WS, " ", proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
				tw.AppendWithReal(pb.TokenType_Punctuation, "(", proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
				for _, n := range spec.Names {
					tw.AppendWithReal(pb.TokenType_WS, "\n\t", proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
					tw.AppendWithReal(pb.TokenType_Identifier, n.Name, proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
					tw.AppendWithReal(pb.TokenType_WS, " …", proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
				}
				tw.AppendWithReal(pb.TokenType_WS, "\n", proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
				tw.AppendWithReal(pb.TokenType_Punctuation, ")", proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
			}
			tw.AppendWithReal(pb.TokenType_WS, "\n\n", proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
		case *ast.TypeSpec:
			tw.AppendWithReal(pb.TokenType_Keyword, kw, proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
			tw.AppendWithReal(pb.TokenType_WS, " ", proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
			tw.AppendWithReal(pb.TokenType_Identifier, spec.Name.Name, proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
			tw.AppendWithReal(pb.TokenType_WS, " …", proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
			tw.AppendWithReal(pb.TokenType_WS, "\n\n", proto.Uint32(off), proto.Uint32(uint32(lineCtr)), href)
		}
	}
}

func UimReferenceContext(decl ast.Node) *string {
	switch d := decl.(type) {
	case *ast.FuncDecl:
		return &d.Name.Name
	case *ast.GenDecl:
		ts := tokStr(d.Tok)
		if len(d.Specs) != 1 {
			return &ts
		}
		switch spec := d.Specs[0].(type) {
		case *ast.ValueSpec:
			if len(spec.Names) != 1 {
				return &ts
			}
			return &spec.Names[0].Name
		case *ast.TypeSpec:
			return &spec.Name.Name
		}

	default:
		return nil
	}

	return nil
}
