package iterator

import (
	"math/rand"
	"sort"
)

var (
	testRandArrays [][]int
)

func init() {
	println("init testRandArrays")
	testRandArrays = [][]int{
		randArray(1000),
		randArray(2000),
		randArray(3000)}
}

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

func arraysCombine(a [][]int, f func(a, b []int) []int) []int {
	size := len(a)
	if size == 0 {
		a = [][]int{[]int{}, []int{}}
	}
	if size == 1 {
		a = append(a, []int{})
	}
	c := f(a[0], a[1])
	for i := 2; i < size; i++ {
		c = f(c, a[i])
	}
	return c
}
