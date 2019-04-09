package main

import (
	"math"
	"runtime"
)

func MultithreadMergeSort(A []int, N int) []int {
	if N <= 0 {
		N = runtime.NumCPU()
	}
	if len(A) <= 1 {
		return A
	}
}

func split(l, r, N int) (indexes []int) {
	indexes = make([]int, l)
	step := int(math.Round(float64(r-l) / float64(N)))
	for i := l; i < r; i += step {
		indexes = append(indexes, i)
	}
	indexes = append(indexes, r)
	return indexes
}

func MergeSort(A []int, l, r int) {
	if l < r {
		m := l + r/2
		MergeSort(A, l, m)
		MergeSort(A, m+1, r)
		Merge
	}
}

func Merge(A []int, indexes []int) {

}
