package main

//go:generate go run github.com/worldiety/choicegen@latest

type None struct{}

type A int
type B int
type C string

type Dec int
type Inc int

type DudeError string
type OrderError string

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
