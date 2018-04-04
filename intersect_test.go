package iterator

import (
	"sort"
	"testing"
)

func intersect(a, b []int) []int {
	asize := len(a)
	bsize := len(b)
	csize := asize
	if csize > bsize {
		csize = bsize
	}
	c := make([]int, csize)
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

func intersectA(a [][]int) []int {
	size := len(a)
	if size < 2 {
		return []int{}
	}
	c := intersect(a[0], a[1])
	for i := 2; i < size; i++ {
		c = intersect(c, a[i])
	}
	return c
}

func TestIntersectIterator(t *testing.T) {
	iter := NewIntersectIterator([]Iterator{
		NewArrayIterator([]int{0, 1, 2, 3}),
		NewArrayIterator([]int{1, 2, 3, 4}),
	})
	v, ok := iter.Next()
	if !ok || v != 1 {
		t.Fail()
	}
	v, ok = iter.Next()
	if !ok || v != 2 {
		t.Fail()
	}
	v, ok = iter.Next()
	if !ok || v != 3 {
		t.Fail()
	}
	v, ok = iter.Next()
	if ok || v != 0 {
		t.Fail()
	}
}

func TestIntersectIteratorIssue0(t *testing.T) {
	iter := NewIntersectIterator([]Iterator{
		NewArrayIterator([]int{4037200794235010051, 6129484611666145821}),
		NewArrayIterator([]int{3916589616287113937, 6334824724549167320}),
	})
	v, ok := iter.Next()
	if ok || v != 0 {
		t.Fail()
	}
}

func TestIntersectIteratorBrut(t *testing.T) {
	for i := 0; i < 1000; i++ {
		it0 := arrayIteratorRnd(i)
		it1 := arrayIteratorRnd(i)
		res := make([]int, 0, i)
		iter := NewIntersectIterator([]Iterator{it0, it1})
		for {
			v, ok := iter.Next()
			if !ok {
				break
			}
			res = append(res, v)
		}
		if !sort.IntsAreSorted(res) {
			t.Fail()
		}
	}
}

func BenchmarkIntersectIterator(b *testing.B) {
	b.StopTimer()
	iter := NewIntersectIterator([]Iterator{arrayIterator0, arrayIterator1})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		iter.Reset()
		for {
			v, ok := iter.Next()
			if !ok {
				break
			}
			_ = v
		}
	}
}
