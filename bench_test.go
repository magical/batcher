package batcher

import (
	"crypto/subtle"
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func BenchmarkSortAscending(b *testing.B) {
	for n := 10; n < 1e6; n *= 10 {
		b.Run("n="+strconv.Itoa(n), func(b *testing.B) {
			arr := make([]int, n)
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				for i := range arr {
					arr[i] = i
				}
				b.StartTimer()
				Sort(IntSlice(arr))
			}
		})
	}
}

func BenchmarkSortRandom(b *testing.B) {
	for n := 10; n < 1e6; n *= 10 {
		b.Run("n="+strconv.Itoa(n), func(b *testing.B) {
			arr := make([]int, n)
			perm := rand.Perm(n)
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(arr, perm)
				b.StartTimer()
				Sort(IntSlice(arr))
			}
		})
	}
}

func BenchmarkConstantTimeSort(b *testing.B) {
	for n := 10; n < 1e6; n *= 10 {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			arr := make([]int32, n)
			perm := rand.Perm(n)
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				for i, v := range perm {
					arr[i] = int32(v)
				}
				b.StartTimer()
				Sort(ctslice(arr))
			}
		})
	}
}

type ctslice []int32

func (s ctslice) Len() int        { return len(s) }
func (s ctslice) Minmax(i, j int) { ctminmax(s, i, j) }

func ctminmax(x []int32, i, j int) {
	a, b := int(x[i]), int(x[j])
	le := subtle.ConstantTimeLessOrEq(a, b)
	x[i] = int32(subtle.ConstantTimeSelect(le, a, b))
	x[j] = int32(subtle.ConstantTimeSelect(1-le, a, b))
}

// for comparison
func BenchmarkStdSortRandom(b *testing.B) {
	for n := 10; n < 1e6; n *= 10 {
		b.Run("n="+strconv.Itoa(n), func(b *testing.B) {
			arr := make([]int, n)
			perm := rand.Perm(n)
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(arr, perm)
				b.StartTimer()
				sort.Sort(sort.IntSlice(arr))
			}
		})
	}
}
