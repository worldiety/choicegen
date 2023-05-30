package main

func intersect[T comparable](a, b []T) []T {
	var intersection []T
	for _, aT := range a {
		inBoth := false
		for _, bT := range b {
			if aT == bT {
				inBoth = true
				break
			}
		}

		if inBoth {
			intersection = append(intersection, aT)
		}
	}

	return intersection
}

func mapEach[T, R any](v []T, f func(T) R) []R {
	var r []R
	for _, t := range v {
		r = append(r, f(t))
	}

	return r
}
