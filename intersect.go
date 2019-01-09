package iterator

type InterIter struct {
	a Iterator
	b Iterator
}

func NewInterIter(a, b Iterator) *InterIter {
	return &InterIter{a: a, b: b}
}

func (it *InterIter) Reset() {
	it.a.Reset()
	it.b.Reset()
}

func (it *InterIter) Next() (int, bool) {
	a, ok := it.a.Next()
	if !ok {
		return 0, false
	}
	b, ok := it.b.Next()
	if !ok {
		return 0, false
	}
	for {
		if a < b {
			if a, ok = it.a.Next(); !ok {
				return 0, false
			}
			continue
		}
		if a > b {
			if b, ok = it.b.Next(); !ok {
				return 0, false
			}
			continue
		}
		break
	}
	return a, true
}
