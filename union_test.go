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
	a := []int{0, 100, 200, 200, 300, 350, 400}
	b := []int{200, 200, 400, 500}
	c := []int{200, 200, 400}
	d := arraysCombine([][]int{a, b, c}, union)
	if !arraysIsEqual([]int{0, 100, 200, 200, 300, 350, 400, 500}, d) {
		t.Fail()
	}
}
