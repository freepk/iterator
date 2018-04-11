package iterator

import (
	"testing"

	"github.com/freepk/arrays"
)

func TestIntersect(t *testing.T) {
	a := []int{0, 100, 200, 300, 350, 400}
	b := []int{200, 400, 500}
	c := []int{200, 400}
	d := arraysCombine([][]int{a, b, c}, arrays.Intersect)
	if !arrays.IsEqual(c, d) {
		t.Fail()
	}
}

func TestIntersectIterator(t *testing.T) {
	a := arraysCombine(testRandArrays, arrays.Intersect)
	b := NewIntersectIterator(arraysToIterators(testRandArrays))
	c := make([]int, 0)
	for {
		v, ok := b.Next()
		if !ok {
			break
		}
		c = append(c, v)
	}
	if !arrays.IsEqual(a, c) {
		t.Fail()
	}

	c = c[:0]
	b.Reset()
	for {
		v, ok := b.Next()
		if !ok {
			break
		}
		c = append(c, v)
	}
	if !arrays.IsEqual(a, c) {
		t.Fail()
	}
}

func BenchmarkIntersect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arraysCombine(testRandArrays, arrays.Intersect)
	}
}

func BenchmarkIntersectIterator(b *testing.B) {
	it := NewIntersectIterator(arraysToIterators(testRandArrays))
	for i := 0; i < b.N; i++ {
		it.Reset()
		for {
			_, ok := it.Next()
			if !ok {
				break
			}
		}
	}
}
