package iterator

import (
	"testing"
)

func union(a, b []int) []int {
	asize := len(a)
	bsize := len(b)
	c := make([]int, 0, asize+bsize)
	i := 0
	j := 0
	for i < asize && j < bsize {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
			continue
		}
		if a[i] > b[j] {
			c = append(c, b[j])
			j++
			continue
		}
		c = append(c, a[i])
		c = append(c, b[j])
		i++
		j++
	}
	for i < asize {
		c = append(c, a[i])
		i++
	}
	for j < bsize {
		c = append(c, b[j])
		j++
	}
	return c
}

func TestUnion(t *testing.T) {
	a := []int{0, 100, 200, 300, 350, 400}
	b := []int{400, 500}
	c := []int{200, 400}
	d := arraysCombine([][]int{a, b, c}, union)
	if !arraysIsEqual([]int{0, 100, 200, 200, 300, 350, 400, 400, 400, 500}, d) {
		t.Fail()
	}
}

func TestUnionIterator(t *testing.T) {
	a := arraysCombine(testRandArrays, union)
	b := NewUnionIterator(arraysToIterators(testRandArrays))
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
}

func BenchmarkUnion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arraysCombine(testRandArrays, union)
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
