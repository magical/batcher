// Package batcher implements a Batcher odd-even sorting network.
//
// Sorting networks are built on an operation called minmax which sorts two elements.
//
// 	func minmax(i, j int) {
// 		if a[i] > a[j] {
// 			a[i], a[j] = a[j], a[i]
// 		}
// 	}
//
// Or, if we imagine a circuit with the array elements flowing from left to right,
// we can draw minmax like this:
//
// 	a[i] >-----+-----> min(a[i], a[j])
// 	           |
// 	           |
// 	a[j] >-----+-----> max(a[i], a[j])
//
// Given a particular array length, it is possible to build a network of
// minmax operations such that the output is always sorted.
// Surprisingly, this means that the order of the minmax operations is
// independent of the array contents (this is not true of most other sorting
// algorithm, like quicksort, merge sort and heap sort to name a few.)
//
// Because the memory access pattern is constant for any given array length,
// this has the nice property that if the implementation of minmax is constant time,
// then the sorting network is constant time as well.
//
// There are many ways of laying out a sorting network.
// Batcher's odd-even sorting network is a fairly efficient one,
// performing (n log₂ n)(log₂ n - 1)/4 + n - 1 minmax operations.
package batcher

// Adapted from slide 41 of
// https://cr.yp.to/talks/2018.07.11/slides-djb-20180711-sorting-a4.pdf

import "reflect"

func sortFunc(n int, minmax func(i, j int)) {
	if n < 2 {
		return
	}

	// find the largest power of two less than n
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
	sortFunc(x.Len(), x.Minmax)
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
	minmax := func(i, j int) {
		if x.Less(j, i) {
			x.Swap(j, i)
		}
	}
	sortFunc(n, minmax)
}

// SortSlice sorts a slice given the provided minmax function.
// If it is not a slice, SortSlice panics.
func SortSlice(slice interface{}, minmax func(i, j int)) {
	v := reflect.ValueOf(slice)
	n := v.Len()
	sortFunc(n, minmax)
}
