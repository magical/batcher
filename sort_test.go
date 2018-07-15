package batcher

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	x := rand.Perm(40)
	Sort(IntSlice(x))
	if !sorted(x) {
		t.Errorf("want sorted, got %v", x)
	}
}

type IntSlice []int

func (s IntSlice) Len() int        { return len(s) }
func (s IntSlice) Minmax(i, j int) { minmax([]int(s), i, j) }

func TestSortLess(t *testing.T) {
	x := rand.Perm(40)
	SortLess(sort.IntSlice(x))
	if !sorted(x) {
		t.Errorf("want sorted, got %v", x)
	}
}

func TestSortSlice(t *testing.T) {
	x := rand.Perm(40)
	SortSlice(x, func(i, j int) { minmax(x, i, j) })
	if !sorted(x) {
		t.Errorf("want sorted, got %v", x)
	}
}

func TestSortSliceLess(t *testing.T) {
	x := rand.Perm(40)
	SortSlice(x, func(i, j int) { minmax(x, i, j) })
	if !sorted(x) {
		t.Errorf("want sorted, got %v", x)
	}
}

func sorted(x []int) bool {
	for i := 0; i < len(x)-1; i++ {
		if !(x[i] < x[i+1]) {
			return false
		}
	}
	return true
}

func minmax(x []int, i, j int) {
	if x[i] > x[j] {
		x[i], x[j] = x[j], x[i]
	}
}
