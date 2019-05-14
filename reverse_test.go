package iterator

import (
	"testing"

	"github.com/freepk/arrays"
)

func TestReversSimpleIter(t *testing.T) {
	a := []int{0, 1}
	it := NewReversIter(NewArrayIter(a))

	elem, _ := it.Next()
	if elem != 1 {
		t.Fail()
	}

	elem, ok := it.Prev()
	if ok {
		t.Fail()
	}

	elem, ok = it.Next()
	elem, ok = it.Next()
	if elem != 0 {
		t.Fail()
	}

	elem, ok = it.Prev()
	if !ok || elem != 1 {
		t.Fail()
	}
}

func TestReverseIter(t *testing.T) {
	a0 := randArray(10)
	it := NewReversIter(NewArrayIter(a0))
	
	out_next := make([]int, 0)
	for {
		if v, ok := it.Next(); !ok {
			break
		} else {
			out_next = append(out_next, v)
		}
	}

	if !arrays.IsEqual(out_next, arrays.Reverse(a0)) {
		t.Fail()
	}

	out_prev := make([]int, 0)
	for {
		if v, ok := it.Prev(); !ok {
			break
		} else {
			out_prev = append(out_prev, v)
		}
	}

	if !arrays.IsEqual(out_prev, a0) {
		t.Fail()
	}
}

func BenchmarkReversIter(b *testing.B) {
	a := randArray(10000)
	it := NewReversIter(NewArrayIter(a))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it.Reset()
		for {
			if _, ok := it.Next(); !ok {
				break
			}
		}
	}
}
