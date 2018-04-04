package iterator

type UnionIterator struct {
        iterators []Iterator
        values    []int
        remains   int
}

func NewUnionIterator(iterators []Iterator) *UnionIterator {
        size := len(iterators)
        iter := &UnionIterator{
                iterators: iterators,
                values:    make([]int, size),
                remains:   size}
        iter.Reset()
        return iter
}

func (it *UnionIterator) swap(offset int) {
        it.remains--
        it.iterators[it.remains], it.iterators[offset] = it.iterators[offset], it.iterators[it.remains]
        it.values[it.remains], it.values[offset] = it.values[offset], it.values[it.remains]
}

func (it *UnionIterator) Reset() {
        ok := false
        it.remains = len(it.iterators)
        for i := 0; i < it.remains; i++ {
                it.iterators[i].Reset()
                it.values[i], ok = it.iterators[i].Next()
                if !ok {
                        it.swap(i)
                }
        }
}

func (it *UnionIterator) Next() (int, bool) {
        if it.remains == 0 {
                return 0, false
        }
        offset := 0
        advice := it.values[0]
        for i := 1; i < it.remains; i++ {
                if advice > it.values[i] {
                        offset = i
                        advice = it.values[i]
                }
        }
        value, ok := it.iterators[offset].Next()
        if ok {
                it.values[offset] = value
        } else {
                it.swap(offset)
        }
        return advice, true
}
