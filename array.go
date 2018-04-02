package iterator

type ArrayIterator struct {
        items  []int
        offset int
}

func NewArrayIterator(items []int) *ArrayIterator {
        return &ArrayIterator{items: items, offset: 0}
}

func (it *ArrayIterator) Reset() {
        it.offset = 0
}

func (it *ArrayIterator) Next() (int, bool) {
        if it.offset < len(it.items) {
                offset := it.offset
                it.offset++
                return it.items[offset], true
        }
        return 0, false
}
