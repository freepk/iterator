package iterator

import (
	"math/rand"
	"sort"

	"github.com/freepk/radix"
)

var (
	arrayIterator0 *ArrayIterator
	arrayIterator1 *ArrayIterator
)

func init() {
	arrayIterator0 = arrayIteratorRnd(30000)
	arrayIterator1 = arrayIteratorRnd(10000)
}

func arrayIteratorRnd(size int) *ArrayIterator {
	if size == 0 {
		return NewArrayIterator([]int{})
	}
	arr := make([]int, size)
	tmp := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Int()
	}
	radix.Ints(arr, tmp, size)
	return NewArrayIterator(arr)
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

func randArrays(a, b int) [][]int {
	n := rand.Intn(a) + 2
	r := make([][]int, n)
	for i := 0; i < n; i++ {
		m := b + rand.Intn(b)
		l := make([]int, m)
		for j := 0; j < m; j++ {
			l[j] = rand.Intn(2 * b)
		}
		sort.Ints(l)
		r[i] = l
	}
	return r
}

func arraysToIterators(a [][]int) []Iterator {
	size := len(a)
	res := make([]Iterator, size)
	for i := 0; i < size; i++ {
		res[i] = NewArrayIterator(a[i])
	}
	return res
}
