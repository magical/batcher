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

func TestNetworkSize(t *testing.T) {
	{
		// log₂ 8 = 3
		const expected = 8*3*(3-1)/4 + 8 - 1
		actual := 0
		sortFunc(8, func(i, j int) {
			actual++
		})
		if actual != expected {
			t.Errorf("128-element array used %d operations, expected %d", actual, expected)
		}
	}

	{
		// log₂ 128 = 7
		const expected = 128*7*(7-1)/4 + 128 - 1
		actual := 0
		sortFunc(128, func(i, j int) {
			actual++
		})
		if actual != expected {
			t.Errorf("128-element array used %d operations, expected %d", actual, expected)
		}
	}
}
