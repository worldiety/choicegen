// Code generated by choice; DO NOT EDIT

package main

import (
	"fmt"
)

// MyGeneratedPageEvents shows the mighty generator.
// In your ears.
//
// Cool.
type MyGeneratedPageEvents interface {
	// Inc | Dec | None

	isMyGeneratedPageEvents() bool // marker method for MyGeneratedPageEvents

}

type MyError2 interface {
	// DudeError | OrderError

	isMyError2() bool // marker method for MyError2

}

type Sum1 interface {
	// A | B | C

	isSum1() bool // marker method for Sum1

}

type Sum2 interface {
	// A

	isSum1() bool // marker method for Sum1

	isSum2() bool // marker method for Sum2

}

// OrderError is one of interface { MyError2 }

func (_ OrderError) isMyError2() bool { return true }

func (o OrderError) Error() string { return fmt.Sprintf("%T: %v", o, o) }

// A is one of interface { Sum1 | Sum2 }

func (_ A) isSum1() bool { return true }

func (_ A) isSum2() bool { return true }

// B is one of interface { Sum1 }

func (_ B) isSum1() bool { return true }

// C is one of interface { Sum1 }

func (_ C) isSum1() bool { return true }

// Inc is one of interface { MyGeneratedPageEvents }

func (_ Inc) isMyGeneratedPageEvents() bool { return true }

// Dec is one of interface { MyGeneratedPageEvents }

func (_ Dec) isMyGeneratedPageEvents() bool { return true }

// None is one of interface { MyGeneratedPageEvents }

func (_ None) isMyGeneratedPageEvents() bool { return true }

// DudeError is one of interface { MyError2 }

func (_ DudeError) isMyError2() bool { return true }

func (d DudeError) Error() string { return fmt.Sprintf("%T: %v", d, d) }

// MatchMyGeneratedPageEvents checks each type case and panics either if choiceType is nil or if an interface compatible
// type has been passed but is not part of the sum type specification.
// Each case must be handled and evaluate properly, so nil functions will panic.
func MatchMyGeneratedPageEvents[R any](choiceType MyGeneratedPageEvents, matchInc func(Inc) R, matchDec func(Dec) R, matchNone func(None) R) R {
	switch t := choiceType.(type) {
	case Inc:
		return matchInc(t)
	case Dec:
		return matchDec(t)
	case None:
		return matchNone(t)
	}

	panic(fmt.Sprintf("%T is not part of the choice type MyGeneratedPageEvents", choiceType))
}

// MatchMyError2 checks each type case and panics either if choiceType is nil or if an interface compatible
// type has been passed but is not part of the sum type specification.
// Each case must be handled and evaluate properly, so nil functions will panic.
func MatchMyError2[R any](choiceType MyError2, matchDudeError func(DudeError) R, matchOrderError func(OrderError) R) R {
	switch t := choiceType.(type) {
	case DudeError:
		return matchDudeError(t)
	case OrderError:
		return matchOrderError(t)
	}

	panic(fmt.Sprintf("%T is not part of the choice type MyError2", choiceType))
}

// MatchSum1 checks each type case and panics either if choiceType is nil or if an interface compatible
// type has been passed but is not part of the sum type specification.
// Each case must be handled and evaluate properly, so nil functions will panic.
func MatchSum1[R any](choiceType Sum1, matchA func(A) R, matchB func(B) R, matchC func(C) R) R {
	switch t := choiceType.(type) {
	case A:
		return matchA(t)
	case B:
		return matchB(t)
	case C:
		return matchC(t)
	}

	panic(fmt.Sprintf("%T is not part of the choice type Sum1", choiceType))
}

// MatchSum2 checks each type case and panics either if choiceType is nil or if an interface compatible
// type has been passed but is not part of the sum type specification.
// Each case must be handled and evaluate properly, so nil functions will panic.
func MatchSum2[R any](choiceType Sum2, matchA func(A) R) R {
	switch t := choiceType.(type) {
	case A:
		return matchA(t)
	}

	panic(fmt.Sprintf("%T is not part of the choice type Sum2", choiceType))
}