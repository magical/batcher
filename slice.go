// +build go18

package batcher

import "reflect"

// SortSlice sorts a slice given the provided less function.
// If it is not a slice, SortSliceLess panics.
func SortSliceLess(slice interface{}, less func(i, j int) bool) {
	v := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)

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
				if less(i+p, i) {
					swap(i+p, i)
				}
			}
		}
		for q := t; q > p; q >>= 1 {
			for i := 0; i < n-q; i++ {
				if i&p == 0 {
					if less(i+p, i+q) {
						swap(i+p, i+q)
					}
				}
			}
		}
	}

}
