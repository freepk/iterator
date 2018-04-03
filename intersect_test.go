package iterator

import (
	"github.com/freepk/radix"
	"math/rand"
	"sort"
	"testing"
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

func TestIntersectIterator(t *testing.T) {
	iter := NewIntersectIterator([]Iterator{
		NewArrayIterator([]int{0, 1, 2, 3}),
		NewArrayIterator([]int{1, 2, 3, 4}),
	})
	v, ok := iter.Next()
	if !ok || v != 1 {
		t.Fail()
	}
	v, ok = iter.Next()
	if !ok || v != 2 {
		t.Fail()
	}
	v, ok = iter.Next()
	if !ok || v != 3 {
		t.Fail()
	}
	v, ok = iter.Next()
	if ok || v != 0 {
		t.Fail()
	}
}

func TestIntersectIteratorIssue0(t *testing.T) {
	iter := NewIntersectIterator([]Iterator{
		NewArrayIterator([]int{4037200794235010051, 6129484611666145821}),
		NewArrayIterator([]int{3916589616287113937, 6334824724549167320}),
	})
	v, ok := iter.Next()
	if ok || v != 0 {
		t.Fail()
	}
}

func TestIntersectIteratorBrut(t *testing.T) {
	for i := 0; i < 1000; i++ {
		it0 := arrayIteratorRnd(i)
		it1 := arrayIteratorRnd(i)
		res := make([]int, 0, i)
		iter := NewIntersectIterator([]Iterator{it0, it1})
		for {
			v, ok := iter.Next()
			if !ok {
				break
			}
			res = append(res, v)
		}
		if !sort.IntsAreSorted(res) {
			t.Fail()
		}
	}
}

func BenchmarkIntersectIterator(b *testing.B) {
	b.StopTimer()
	iter := NewIntersectIterator([]Iterator{arrayIterator0, arrayIterator1})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		iter.Reset()
		for {
			v, ok := iter.Next()
			if !ok {
				break
			}
			_ = v
		}
	}
}
