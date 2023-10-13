package main

//go:generate go run github.com/worldiety/choicegen@latest

type None struct{}

type A int
type B int
type C string

type Dec int
type Inc int

type Krankheit int
type Urlaub int

type Elternzeit int
type UnbezahlterUrlaub int

type Anstellung int
type Selbstständigkeit int
type Freiwillig int
type Pflicht int

type DudeError string

func (d DudeError) String() any {
	return string(d)
}

func (d DudeError) MyCustomMethod() {

}

func (d DudeError) MyCustom2(a string, x, y int) (bool, error) {
	return false, nil
}

type OrderError string

func (o OrderError) String() any {
	return string(o)
}

func (o OrderError) MyCustomMethod() {

}

func (d OrderError) MyCustom2(a string, x, y int) (bool, error) {
	return false, nil
}

func main() {
	var events MyGeneratedPageEvents

	events = Inc(1)
	events = Dec(2)
	events = None{}
	// events = "hello" doesn't compile

	err := MatchMyGeneratedPageEvents(events,
		func(inc Inc) MyError2 {
			return nil
		}, func(dec Dec) MyError2 {
			return nil
		}, func(none None) MyError2 {
			return nil
		},
	)

	if err != nil {
		panic(err)
	}
}
