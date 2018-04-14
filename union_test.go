package iterator

import (
	"testing"

	"github.com/freepk/arrays"
)

func TestUnion(t *testing.T) {
	a := []int{0, 100, 200, 300, 350, 400}
	b := []int{400, 500}
	c := []int{200, 400}
	d := arrays.Combine([][]int{a, b, c}, arrays.Union)
	if !arrays.IsEqual([]int{0, 100, 200, 200, 300, 350, 400, 400, 400, 500}, d) {
		t.Fail()
	}
}

func TestUnionIterator(t *testing.T) {
	a := arrays.Combine(testRandArrays, arrays.Union)
	b := NewUnionIterator(arraysToIterators(testRandArrays))
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

func BenchmarkUnion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrays.Combine(testRandArrays, arrays.Union)
	}
}

func BenchmarkUnionIterator(b *testing.B) {
	it := NewUnionIterator(arraysToIterators(testRandArrays))
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
