// Package batcher implements a Batcher even-odd sorting network.
package batcher

// Adapted from slide 41 of
// https://cr.yp.to/talks/2018.07.11/slides-djb-20180711-sorting-a4.pdf

import "reflect"

// Interface is the interface required by Sort.
type Interface interface {
	Len() int

	// Minmax compares the elements and positions i and j,
	// and swaps them if element i is greater than element j.
	// In other words it it assignes the min of the two elements to i
	// and the max to j.
	Minmax(i, j int)
}

// Sort sorts x using (n log^2 n)/4 operations.
func Sort(x Interface) {
	n := x.Len()
	if n < 2 {
		return
	}

	t := 1
	for t < n-t {
		t += t
	}
	for p := t; p > 0; p >>= 1 {
		for i := 0; i < n-p; i++ {
			if i&p == 0 {
				x.Minmax(i, i+p)
			}
		}
		for q := t; q > p; q >>= 1 {
			for i := 0; i < n-q; i++ {
				if i&p == 0 {
					x.Minmax(i+p, i+q)
				}
			}
		}
	}
}

type LessSwapper interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// Sort sorts x using (n log^2 n)/4 operations.
// It accepts the same interface as sort.Sort from the standard library.
func SortLess(x LessSwapper) {
	n := x.Len()
	if n < 2 {
		return
	}

	t := 1
	for t < n-t {
		t += t
	}
	for p := t; p > 0; p >>= 1 {
		for i := 0; i < n-p; i++ {
			if i&p == 0 {
				if x.Less(i+p, i) {
					x.Swap(i+p, i)
				}
			}
		}
		for q := t; q > p; q >>= 1 {
			for i := 0; i < n-q; i++ {
				if i&p == 0 {
					if x.Less(i+q, i+p) {
						x.Swap(i+q, i+p)
					}
				}
			}
		}
	}
}

// SortSlice sorts a slice given the provided minmax function.
// If it is not a slice, SortSlice panics.
func SortSlice(slice interface{}, minmax func(i, j int)) {
	v := reflect.ValueOf(slice)
	n := v.Len()
	if n < 2 {
		return
	}

	t := 1
	for t < n-t {
		t += t
	}
	for p := t; p > 0; p >>= 1 {
		for i := 0; i < n-p; i++ {
			if i&p == 0 {
				minmax(i, i+p)
			}
		}
		for q := t; q > p; q >>= 1 {
			for i := 0; i < n-q; i++ {
				if i&p == 0 {
					minmax(i+p, i+q)
				}
			}
		}
	}
}
