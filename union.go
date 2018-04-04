package iterator

type UnionIterator struct {
	iterators []Iterator
}

func NewUnionIterator(iterators []Iterator) *UnionIterator {
	return &UnionIterator{iterators: iterators}
}

func (it *UnionIterator) Reset() {
}

func (it *UnionIterator) Next() (int, bool) {
	return 0, false
}
