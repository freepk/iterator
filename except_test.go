package iterator

import (
	"testing"

	"github.com/freepk/arrays"
)

func TestExcept(t *testing.T) {
	a := []int{0, 100, 100, 200, 300, 350, 400}
	b := []int{400, 500}
	c := []int{200, 400}
	d := combineArrays([][]int{a, b, c}, arrays.Except)
	if !arrays.IsEqual([]int{0, 100, 100, 300, 350}, d) {
		t.Fail()
	}
}

func TestExceptIterator(t *testing.T) {
	a := combineArrays(testRandArrays, arrays.Except)
	b := NewExceptIterator(
		NewArrayIterator(testRandArrays[0]),
		arraysToIterators(testRandArrays[1:]))
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

func TestExceptIteratorX(t *testing.T) {
	a := combineArrays(testRandArrays, arrays.Except)
	b := NewExceptIterator(
		NewArrayIterator(testRandArrays[0]),
		[]Iterator{NewArrUnionIterator(testRandArrays[1:])})
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

func BenchmarkExcept(b *testing.B) {
	for i := 0; i < b.N; i++ {
		combineArrays(testRandArrays, arrays.Except)
	}
}

func BenchmarkExceptIterator(b *testing.B) {
	it := NewExceptIterator(
		NewArrayIterator(testRandArrays[0]),
		arraysToIterators(testRandArrays[1:]))
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

func BenchmarkExceptIteratorX(b *testing.B) {
	it := NewExceptIterator(
		NewArrayIterator(testRandArrays[0]),
		[]Iterator{NewArrUnionIterator(testRandArrays[1:])})
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
