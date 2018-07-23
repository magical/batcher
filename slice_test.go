package batcher

import (
	"math/rand"
	"testing"
)

func TestSortSliceLess(t *testing.T) {
	x := rand.Perm(40)
	SortSliceLess(x, func(i, j int) bool { return x[i] < x[j] })
	if !sorted(x) {
		t.Errorf("want sorted, got %v", x)
	}
}
