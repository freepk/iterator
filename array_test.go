package iterator

import (
	"testing"
)

func TestArrayIterator(t *testing.T) {
	iter := NewArrayIterator([]int{1, 2, 3})
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

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, arr := range testRandArrays {
			for _, v := range arr {
				_ = v
			}
		}
	}
}

func BenchmarkArrayIterator(b *testing.B) {
	iters := arraysToIterators(testRandArrays)
	for i := 0; i < b.N; i++ {
		for _, iter := range iters {
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
}
