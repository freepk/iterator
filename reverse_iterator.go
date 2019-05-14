package iterator

type ReversIter struct {
	a  Iterator
	a0 int
	a1 bool
}

func NewReversIter(a Iterator) *ReversIter {
	it := &ReversIter{a: a}
	it.a.ResetToEnd()
	return it
}

func (it *ReversIter) Reset() {
	it.a.ResetToEnd()
}

func (it *ReversIter) ResetToEnd() {
	it.a.Reset()
}

func (it *ReversIter) Next() (int, bool) {
	return it.a.Prev()
}

func (it *ReversIter) Prev() (int, bool) {
	return it.a.Next()
}
