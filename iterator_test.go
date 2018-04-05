package iterator

import (
	"math/rand"
	"sort"
	"testing"
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

func BenchmarkComplexIterator(z *testing.B) {
	z.StopTimer()
	a := NewArrayIterator([]int{1, 2, 3, 4, 5})
	b := NewArrayIterator([]int{6, 7, 8, 9, 10})
	c := NewArrayIterator([]int{1, 2, 3, 4, 5})
	d := NewArrayIterator([]int{10, 20, 30, 40, 50})
	e := NewUnionIterator([]Iterator{a, b})
	f := NewUnionIterator([]Iterator{c, d})
	g := NewIntersectIterator([]Iterator{e, f})
	h := NewArrayIterator([]int{300, 400, 500})
	r := NewUnionIterator([]Iterator{h, g})
	z.StartTimer()
	for i := 0; i < z.N; i++ {
		r.Reset()
		for {
			_, ok := r.Next()
			if !ok {
				break
			}
		}
	}
}
