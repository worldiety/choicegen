//go:build choice

package main

// MyGeneratedPageEvents shows the mighty generator.
// In your ears.
//
// Cool.
type MyGeneratedPageEvents interface {
	Inc | Dec | None
}

type MyError2 interface {
	error
	DudeError | OrderError
}

type Sum1 interface {
	A | B | C
}

type Sum2 interface {
	A
}
