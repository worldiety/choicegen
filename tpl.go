package main

import (
	_ "embed"
	"fmt"
	"go/format"
	"strings"
	"text/template"
)

//go:embed choice.gotpl
var choiceTplText string

type ChoiceGoTpl struct {
	Package     string
	ChoiceTypes []ChoiceSubSetType
	MemberTypes []ChoiceTypeMember
}

func (c ChoiceGoTpl) HasFmt() bool {
	/*for _, choiceType := range c.ChoiceTypes {
		if choiceType.Error {
			return true
		}
	}

	return false*/
	return true // matchFunc always has default panic
}

func formatSrc(s string) string {
	buf, err := format.Source([]byte(s))
	if err != nil {
		panic(fmt.Sprintf("%s\n%v", buf, err))
	}

	return string(buf)
}

func applyTpl(tpl *template.Template, model ChoiceGoTpl) string {
	var sb strings.Builder
	if err := tpl.Execute(&sb, model); err != nil {
		panic(err)
	}

	return sb.String()
}

func makeSrc(pkg Identifier, types []ChoiceSubSetType) string {
	tpl := template.New("choice.gotpl")
	tpl.Funcs(map[string]any{
		"makeComment": func(s Comment) string {
			tmp := ""
			for _, line := range strings.Split(string(s), "\n") {
				tmp += "//" + line + "\n"
			}

			return tmp
		},

		"joinIdents": func(idents []Identifier) string {
			tmp := ""
			for i, ident := range idents {
				tmp += string(ident)
				if i < len(idents)-1 {
					tmp += " | "
				}
			}
			return tmp
		},
	})

	tpl = template.Must(tpl.Parse(choiceTplText))
	return formatSrc(applyTpl(tpl, ChoiceGoTpl{
		Package:     string(pkg),
		ChoiceTypes: types,
		MemberTypes: determineChoiceTypeMembers(types),
	}))
}
