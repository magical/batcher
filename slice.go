// +build go1.8

package batcher

import "reflect"

// SortSlice sorts a slice given the provided less function.
// If it is not a slice, SortSliceLess panics.
func SortSliceLess(slice interface{}, less func(i, j int) bool) {
	v := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	minmax := func(i, j int) {
		if less(j, i) {
			swap(j, i)
		}
	}
	sortFunc(v.Len(), minmax)
}
