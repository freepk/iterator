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
	for i := 10; i < 100; i++ {
		a := randArrays(i, i*5)
		b := arraysCombine(a, intersect)
		c := NewIntersectIterator(arraysToIterators(a))
		d := make([]int, 0)
		for {
			v, ok := c.Next()
			if !ok {
				break
			}
			d = append(d, v)
		}
		if !arraysIsEqual(b, d) {
			t.Logf("%#v", b)
			t.Logf("%#v", d)
			t.Fail()
		}
	}
}
