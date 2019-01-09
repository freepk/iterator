package iterator

import (
	"testing"
)

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
