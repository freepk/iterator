package iterator

import (
	"sort"
	"testing"
)

func TestUnionIterator(t *testing.T) {
	a := NewArrayIterator([]int{1, 3})
	b := NewArrayIterator([]int{2})
	iter := NewUnionIterator([]Iterator{a, b})
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

func TestUnionIteratorBrut(t *testing.T) {
	for i := 0; i < 1000; i++ {
		it0 := arrayIteratorRnd(i)
		it1 := arrayIteratorRnd(i)
		res := make([]int, 0, i)
		iter := NewUnionIterator([]Iterator{it0, it1})
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

func BenchmarkUnionIterator(b *testing.B) {
	b.StopTimer()
	iter := NewUnionIterator([]Iterator{arrayIterator0, arrayIterator1})
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
