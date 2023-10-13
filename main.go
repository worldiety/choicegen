package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir, _ := os.Getwd()

	fmt.Println("gen choices ", dir)
	pkgname, choiceTypes := parse(dir)
	// hier ansetzen. Z.B. EliminateChoiceTypeInterfaces.
	eliminateChoiceTypeInterfaces()
	superSetTypes := determineSuperSets(choiceTypes)
	src := makeSrc(Identifier(pkgname), superSetTypes)
	if err := os.WriteFile(filepath.Join(dir, "choicetypes.gen.go"), []byte(src), os.ModePerm); err != nil {
		panic(err)
	}
}
