package iterator

import (
	"testing"

	"github.com/freepk/arrays"
)

func TestUnionAll(t *testing.T) {
	a := []int{0, 100, 200, 300, 350, 400}
	b := []int{400, 500}
	c := []int{200, 400}
	d := combineArrays([][]int{a, b, c}, arrays.UnionAll)
	if !arrays.IsEqual([]int{0, 100, 200, 200, 300, 350, 400, 400, 400, 500}, d) {
		t.Fail()
	}
}

func TestUnionAllIterator(t *testing.T) {
	a := combineArrays(testRandArrays, arrays.UnionAll)
	b := NewUnionAllIterator(arraysToIterators(testRandArrays))
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

func TestArrUnionAllIterator(t *testing.T) {
	a := combineArrays(testRandArrays, arrays.UnionAll)
	b := NewArrUnionAllIterator(testRandArrays)
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

func BenchmarkUnionAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		combineArrays(testRandArrays, arrays.UnionAll)
	}
}

func BenchmarkUnionAllIterator(b *testing.B) {
	it := NewUnionAllIterator(arraysToIterators(testRandArrays))
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

func BenchmarkArrUnionAllIterator(b *testing.B) {
	it := NewArrUnionAllIterator(testRandArrays)
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
