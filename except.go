package iterator

type ExceptIter struct {
	a  Iterator
	b  Iterator
	b0 int
	b1 bool
}

func NewExceptIter(a, b Iterator) *ExceptIter {
	it := &ExceptIter{a: a, b: b}
	it.b0, it.b1 = it.b.Next()
	return it
}

func (it *ExceptIter) Reset() {
	it.a.Reset()
	it.b.Reset()
	it.b0, it.b1 = it.b.Next()
}

func (it *ExceptIter) ResetToEnd() {
	it.a.ResetToEnd()
	it.b.ResetToEnd()
	it.b0, it.b1 = it.b.Prev()
}

func (it *ExceptIter) Next() (int, bool) {
	a, ok := it.a.Next()
	if !ok {
		return 0, false
	}
	if !it.b1 {
		return a, true
	}
	for {
		if a == it.b0 {
			if a, ok = it.a.Next(); !ok {
				return 0, false
			}
			continue
		}
		if a > it.b0 {
			if it.b0, it.b1 = it.b.Next(); !it.b1 {
				return a, true
			}
			continue
		}
		break
	}
	return a, true
}

func (it *ExceptIter) Prev() (int, bool) {
	a, ok := it.a.Prev()
	if !ok {
		return 0, false
	}
	if !it.b1 {
		return a, true
	}
	for {
		if a == it.b0 {
			if a, ok = it.a.Prev(); !ok {
				return 0, false
			}
			continue
		}
		if a < it.b0 {
			if it.b0, it.b1 = it.b.Prev(); !it.b1 {
				return a, true
			}
			continue
		}
		break
	}
	return a, true
}
