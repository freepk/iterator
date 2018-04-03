package iterator

import (
	"math/rand"

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
