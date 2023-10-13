package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
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
						ut := reflect.TypeOf(choiceType.Name)
						fmt.Println(ut.Kind())

						if iface, ok := spec.Type.(*ast.InterfaceType); ok {
							for _, field := range iface.Methods.List {
								if ftype, ok := field.Type.(*ast.FuncType); ok {
									fname := field.Names[0].Name // you cannot express a fn without a name

									fn := FuncSpec{
										Name:   Identifier(fname),
										Params: types(ftype.Params.List),
									}

									if ftype.Results != nil {
										fn.Results = types(ftype.Results.List)
									}

									choiceType.Funcs = append(choiceType.Funcs, fn)
									continue
								}
								//fmt.Printf("%s: Type: %#v\n", choiceType.Name, field.Type)
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

	// ergänzt die ChoiceTypes die eingbettet sind um das Interface, in die sie eingebettet wurden. Dadurch wird verhindert,
	// dass sie weitere Methoden benötigen, um das Interface zu erfüllen.
	for _, re := range res {
		for _, choice := range re.Choices { // choices sind die Auswahlmöglichkeiten in den Interfaces
			choiceIsOtherChoice := false
			for i, declaration := range res {
				if declaration.Name == choice {
					choiceIsOtherChoice = true
				}
				if choiceIsOtherChoice {
					declaration.EmbeddedInterfaces = append(declaration.EmbeddedInterfaces, re.Name)
					choiceIsOtherChoice = false
					res[i] = declaration

				}
			}
		}
	}

	return pkgname, res
}

//TODO: Mit mehreren verschachtelten Interfaces testen. func (_ Krankheit) isFehlzeit() bool {return true} hinterlegen lassen

func types(fields []*ast.Field) []Identifier {
	var res []Identifier
	for _, field := range fields {
		if ident, ok := field.Type.(*ast.Ident); ok {
			if len(field.Names) > 0 {
				for range field.Names {
					res = append(res, Identifier(ident.Name))
				}
			} else {
				res = append(res, Identifier(ident.Name))
			}
		} else {
			fmt.Printf("fix me: unsupported ast import: %#v\n", field.Type)
		}

	}
	return res
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

func eliminateChoiceTypeInterfaces() {

}
