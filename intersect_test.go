package iterator

import (
	"testing"
)

func intersect(a, b []int) []int {
	asize := len(a)
	bsize := len(b)
	csize := asize
	if csize > bsize {
		csize = bsize
	}
	c := make([]int, 0, csize)
	i := 0
	j := 0
	for i < asize && j < bsize {
		if a[i] < b[j] {
			i++
			continue
		}
		if a[i] > b[j] {
			j++
			continue
		}
		c = append(c, a[i])
		i++
		j++
	}
	return c
}

func TestIntersect(t *testing.T) {
	a := []int{0, 100, 200, 300, 350, 400}
	b := []int{200, 400, 500}
	c := []int{200, 400}
	d := arraysCombine([][]int{a, b, c}, intersect)
	if !arraysIsEqual(c, d) {
		t.Fail()
	}
}

func TestIntersectIterator(t *testing.T) {
	a := arraysCombine(testRandArrays, intersect)
	b := NewIntersectIterator(arraysToIterators(testRandArrays))
	c := make([]int, 0)
	for {
		v, ok := b.Next()
		if !ok {
			break
		}
		c = append(c, v)
	}
	if !arraysIsEqual(a, c) {
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
	if !arraysIsEqual(a, c) {
		t.Fail()
	}
}

func BenchmarkIntersect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arraysCombine(testRandArrays, intersect)
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
