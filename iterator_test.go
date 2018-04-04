package iterator

import (
	"math/rand"
	"sort"
)

func randArray(size int) []int {
	if size == 0 {
		return []int{}
	}
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = rand.Intn(size * 2)
	}
	sort.Ints(res)
	return res
}

func randArrays(arraysNum, arraySize int) [][]int {
	res := make([][]int, arraysNum)
	for i := 0; i < arraysNum; i++ {
		res[i] = randArray(arraySize)
	}
	return res
}

func arraysIsEqual(a, b []int) bool {
	size := len(a)
	if size != len(b) {
		return false
	}
	for i := 0; i < size; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func arraysToIterators(a [][]int) []Iterator {
	size := len(a)
	res := make([]Iterator, size)
	for i := 0; i < size; i++ {
		res[i] = NewArrayIterator(a[i])
	}
	return res
}
