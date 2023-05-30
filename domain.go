package main

import (
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"strings"
)

// An Identifier follows the rules of a Go identifier, e.g. UnvalidatedOrder.
type Identifier string

// ChoiceName is a unique identifier of an existing local type (e.g. UnvalidatedOrder or ValidatedOrder).
// This type will later receive marker methods.
type ChoiceName Identifier

// Comment may be empty, a single line or multiline text.
type Comment string

// A ChoiceTypeDeclaration aggregates all information about a new ChoiceType which must be generated.
type ChoiceTypeDeclaration struct {
	Doc     Comment
	Name    Identifier
	Choices []Identifier
	Error   bool
}

// SuperSet is a set of local type identifiers which are equal or the super set of something else.
type SuperSet []Identifier

// ChoiceSubSetType annotates all super set types.
type ChoiceSubSetType struct {
	ChoiceTypeDeclaration
	SubsetOf SuperSet
}

// DetermineSuperSets transforms from declarations to subset types.
type DetermineSuperSets func([]ChoiceTypeDeclaration) []ChoiceSubSetType

// A ChoiceTypeMember belongs to all its contained ChoiceTypes.
type ChoiceTypeMember struct {
	Name      Identifier
	BelongsTo []Identifier
	Error     bool
}

func (c ChoiceTypeMember) ShortName() string {
	return strings.ToLower(string(rune(c.Name[0])))
}

func determineSuperSets(decls []ChoiceTypeDeclaration) []ChoiceSubSetType {
	var types []ChoiceSubSetType
	for _, decl := range decls {
		types = append(types, ChoiceSubSetType{
			ChoiceTypeDeclaration: decl,
			SubsetOf: mapEach(superSets(decl, decls), func(t ChoiceTypeDeclaration) Identifier {
				return t.Name
			}),
		})
	}

	return types
}

func determineChoiceTypeMembers(decls []ChoiceSubSetType) []ChoiceTypeMember {
	members := map[Identifier]ChoiceTypeMember{}
	for _, decl := range decls {
		for _, choice := range decl.Choices {
			t := members[choice]
			t.Name = choice
			t.Error = decl.Error
			if !slices.Contains(t.BelongsTo, decl.Name) {
				t.BelongsTo = append(t.BelongsTo, decl.Name)
			}
			members[choice] = t
		}
	}

	return maps.Values(members)
}

func superSets(decls ChoiceTypeDeclaration, other []ChoiceTypeDeclaration) []ChoiceTypeDeclaration {
	var superSets []ChoiceTypeDeclaration
	dstSet := slices.Clone(decls.Choices)
	for _, declaration := range other {
		otherDecl := slices.Clone(declaration.Choices)
		slices.Sort(otherDecl)
		if slices.Equal(intersect(dstSet, otherDecl), dstSet) {
			superSets = append(superSets, declaration)
		}
	}

	return superSets
}
