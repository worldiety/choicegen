package main

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func parse(dir string) (pkgname string, _ []ChoiceTypeDeclaration) {
	var res []ChoiceTypeDeclaration
	fileset := token.NewFileSet()
	pkg, err := parser.ParseDir(fileset, dir, func(info fs.FileInfo) bool {
		if !strings.HasSuffix(info.Name(), ".go") {
			return false
		}

		var tmp [1024]byte
		f, err := os.Open(filepath.Join(dir, info.Name()))
		if err != nil {
			panic(err)
		}
		defer f.Close()

		for i := range tmp {
			tmp[i] = 0
		}
		_, err = f.Read(tmp[:])
		if err != nil && err != io.EOF {
			panic(err)
		}

		if bytes.Contains(tmp[:], []byte("go:build choice")) {
			return true
		}

		return false
	}, parser.ParseComments)

	if err != nil {
		panic(err)
	}

	for _, p := range pkg {
		pkgname = p.Name
		ast.Walk(VisitorFunc(func(node ast.Node) bool {
			var comment string
			if genDecl, ok := node.(*ast.GenDecl); ok {
				comment = genDecl.Doc.Text()

				for _, spec := range genDecl.Specs {
					if spec, ok := spec.(*ast.TypeSpec); ok {
						choiceType := ChoiceTypeDeclaration{
							Name: Identifier(spec.Name.Name),
							Doc:  Comment(comment),
						}

						if iface, ok := spec.Type.(*ast.InterfaceType); ok {
							for _, field := range iface.Methods.List {
								idents := extractIdentifiers(field)
								if len(idents) == 1 && idents[0] == "error" {
									choiceType.Error = true
								} else {
									choiceType.Choices = idents
								}
							}
							res = append(res, choiceType)

						}
					}
				}
			}

			return true
		}), p)

	}

	return pkgname, res
}

func extractIdentifiers(field *ast.Field) []Identifier {
	var res []Identifier
	ast.Walk(VisitorFunc(func(node ast.Node) bool {
		if ident, ok := node.(*ast.Ident); ok {
			res = append(res, Identifier(ident.Name))
		}

		return true
	}), field)

	return res
}

type VisitorFunc func(node ast.Node) bool

func (v VisitorFunc) Visit(node ast.Node) (w ast.Visitor) {
	if v(node) {
		return v
	}

	return nil
}
