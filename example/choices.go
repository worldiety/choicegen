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
	MyCustomMethod()
	MyCustom2(a string, x, y int) (bool, error)
}

type Sum1 interface {
	A | B | C
}

type Sum2 interface {
	A
}

type Ausfall interface {
	UnbezahlteFehlzeit
}

type ÜberFehlzeit interface {
	Fehlzeit
}

type Fehlzeit interface {
	BezahlteFehlzeit | UnbezahlteFehlzeit
}

type BezahlteFehlzeit interface {
	Krankheit
}

type UnbezahlteFehlzeit interface {
	Elternzeit
}

type Arbeit interface {
	BezahlteArbeit | UnbezahlteArbeit | SonstigeArbeit
}

type BezahlteArbeit interface {
	Anstellung
}

type UnbezahlteArbeit interface {
	Selbstständigkeit
}

type SonstigeArbeit interface {
	Praktikum | Hospitation
}

type Praktikum interface {
	Pflicht | Freiwillig
}

type Hospitation interface {
	Pflicht | Freiwillig
}
