package iterator

import (
        "testing"
)

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
