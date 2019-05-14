package iterator

import (
	"testing"

	"github.com/freepk/arrays"
)

func TestExcept(t *testing.T) {
	a0 := randArray(10)
	a1 := randArray(5)
	it := NewExceptIter(NewArrayIter(a0), NewArrayIter(a1))
	it.Reset()
	out := make([]int, 0)
	for {
		if v, ok := it.Next(); !ok {
			break
		} else {
			out = append(out, v)
		}
	}
	if !arrays.IsEqual(out, arrays.Except(a0, a1)) {
		t.Fail()
	}
}

func TestExceptPrev(t *testing.T) {
	a0 := randArray(10)
	a1 := randArray(5)
	it := NewExceptIter(NewArrayIter(a0), NewArrayIter(a1))
	it.ResetToEnd()
	out := make([]int, 0)
	for {
		if v, ok := it.Prev(); !ok {
			break
		} else {
			out = append(out, v)
		}
	}

	if !arrays.IsEqual(out, arrays.Reverse(arrays.Except(a0, a1))) {
		t.Fail()
	}
}

func BenchmarkExcept(b *testing.B) {
	a0 := randArray(3000)
	a1 := randArray(2000)
	a2 := randArray(1000)
	it := NewExceptIter(NewExceptIter(NewArrayIter(a0), NewArrayIter(a1)), NewArrayIter(a2))
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

func BenchmarkExceptPrev(b *testing.B) {
	a0 := randArray(3000)
	a1 := randArray(2000)
	a2 := randArray(1000)
	it := NewExceptIter(NewExceptIter(NewArrayIter(a0), NewArrayIter(a1)), NewArrayIter(a2))
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

