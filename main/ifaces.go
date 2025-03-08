package main

import (
	"go/ast"
	"go/types"
)

func detectInterfaceImplementations(info *types.Info, files []*ast.File) map[types.Type][]types.Type {
	implementations := make(map[types.Type][]types.Type)

	for _, file := range files {
		var enclosingFunc *types.Signature

		ast.Inspect(file, func(n ast.Node) bool {
			switch node := n.(type) {
			case *ast.FuncDecl:
				// Track the enclosing function's signature
				if obj := info.Defs[node.Name]; obj != nil {
					if fn, ok := obj.(*types.Func); ok {
						enclosingFunc = fn.Type().(*types.Signature)
					}
				}
				return true

			case *ast.AssignStmt:
				// Check assignments to interface variables
				for i, rhs := range node.Rhs {
					if i >= len(node.Lhs) {
						break
					}
					if tv, ok := info.Types[node.Lhs[i]]; ok {
						lhsType := tv.Type
						if rhsTv, ok := info.Types[rhs]; ok {
							rhsType := rhsTv.Type
							if isInterface(lhsType) && types.AssignableTo(rhsType, lhsType) {
								addImplementation(implementations, lhsType, rhsType)
							}
						}
					}
				}

			case *ast.CallExpr:
				// Check type conversions and function calls with interface parameters
				if tv, ok := info.Types[node.Fun]; ok {
					if sig, ok := tv.Type.(*types.Signature); ok {
						// Check function parameters
						params := sig.Params()
						for i := 0; i < params.Len() && i < len(node.Args); i++ {
							paramType := params.At(i).Type()
							if argTv, ok := info.Types[node.Args[i]]; ok {
								argType := argTv.Type
								if isInterface(paramType) && types.AssignableTo(argType, paramType) {
									addImplementation(implementations, paramType, argType)
								}
							}
						}
					}
				}

			case *ast.ReturnStmt:
				// Check return statements against enclosing function's signature
				if enclosingFunc != nil {
					results := enclosingFunc.Results()
					for i, result := range node.Results {
						if i >= results.Len() {
							break
						}
						expectedType := results.At(i).Type()
						if tv, ok := info.Types[result]; ok {
							actualType := tv.Type
							if isInterface(expectedType) && types.AssignableTo(actualType, expectedType) {
								addImplementation(implementations, expectedType, actualType)
							}
						}
					}
				}
			}
			return true
		})
	}

	return implementations
}

func isInterface(t types.Type) bool {
	if named, ok := t.(*types.Named); ok {
		t = named.Underlying()
	}
	_, ok := t.(*types.Interface)
	return ok
}

func addImplementation(implementations map[types.Type][]types.Type, iface, impl types.Type) {
	// Normalize to get the underlying named type if possible
	if named, ok := impl.(*types.Named); ok {
		impl = named.Underlying()
	}

	// Check if we already recorded this implementation
	for _, existing := range implementations[iface] {
		if types.Identical(existing, impl) {
			return
		}
	}

	implementations[iface] = append(implementations[iface], impl)
}
