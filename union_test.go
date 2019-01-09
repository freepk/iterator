package iterator

import (
	"testing"

	"github.com/freepk/arrays"
)

func TestUnionIter(t *testing.T) {
	return
	a0 := randArray(3000)
	a1 := randArray(2000)
	it := NewUnionIter(NewArrayIter(a0), NewArrayIter(a1))
	it.Reset()
	out := make([]int, 0)
	for {
		if v, ok := it.Next(); !ok {
			break
		} else {
			out = append(out, v)
		}
	}
	if !arrays.IsEqual(out, arrays.Union(a0, a1)) {
		t.Fail()
	}
}

func BenchmarkUnion(b *testing.B) {
	a0 := randArray(3000)
	a1 := randArray(2000)
	a2 := randArray(1000)
	it := NewUnionIter(NewUnionIter(NewArrayIter(a0), NewArrayIter(a1)), NewArrayIter(a2))
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
