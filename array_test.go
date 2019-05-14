package iterator

import (
	"testing"
)

func TestArraySimpleIter(t *testing.T) {
	a := []int{0, 1}
	it := NewArrayIter(a)

	elem, _ := it.Next()
	if elem != 0 {
		t.Fail()
	}

	elem, ok := it.Prev()
	if ok {
		t.Fail()
	}

	elem, ok = it.Next()
	elem, ok = it.Next()
	if elem != 1 {
		t.Fail()
	}

	elem, ok = it.Prev()
	if !ok || elem != 0 {
		t.Fail()
	}
}

func TestArrayIter(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	it := NewArrayIter(a)
	it.Reset()
	out := make([]int, 0)
	for {
		if v, ok := it.Next(); !ok {
			break
		} else {
			out = append(out, v)
		}
	}
	if !isEqual(out, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Fail()
	}
}


func TestArrayIterPrev(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	it := NewArrayIter(a)
	it.ResetToEnd()
	out := make([]int, 0)
	for {
		if v, ok := it.Prev(); !ok {
			break
		} else {
			out = append(out, v)
		}
	}
	if !isEqual(out, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}) {
		t.Fail()
	}
}

func BenchmarkArrayIter(b *testing.B) {
	a := randArray(10000)
	it := NewArrayIter(a)
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

func BenchmarkArrayIterPrev(b *testing.B) {
	a := randArray(10000)
	it := NewArrayIter(a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it.ResetToEnd()
		for {
			if _, ok := it.Prev(); !ok {
				break
			}
		}
	}
}
