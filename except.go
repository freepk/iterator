package iterator

type ExceptIterator struct {
	a []Iterator
	v []int
	n int
}

func NewExceptIterator(a []Iterator) *ExceptIterator {
	return nil
}

func (it *ExceptIterator) Reset() {
}

func (it *ExceptIterator) Next() (int, bool) {
	return 0, false
}
