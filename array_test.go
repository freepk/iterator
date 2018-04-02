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
