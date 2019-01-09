package iterator

type UnionAllIter struct {
	a  Iterator
	b  Iterator
	a0 int
	a1 bool
	b0 int
	b1 bool
}

func NewUnionAllIter(a, b Iterator) *UnionAllIter {
	it := &UnionAllIter{a: a, b: b}
	it.a0, it.a1 = it.a.Next()
	it.b0, it.b1 = it.b.Next()
	return it
}

func (it *UnionAllIter) Reset() {
	it.a.Reset()
	it.b.Reset()
	it.a0, it.a1 = it.a.Next()
	it.b0, it.b1 = it.b.Next()
}

func (it *UnionAllIter) Next() (int, bool) {
	if it.a1 && it.b1 {
		if it.a0 < it.b0 {
			v := it.a0
			it.a0, it.a1 = it.a.Next()
			return v, true
		}
		v := it.b0
		it.b0, it.b1 = it.b.Next()
		return v, true
	} else if it.a1 {
		v := it.a0
		it.a0, it.a1 = it.a.Next()
		return v, true
	} else if it.b1 {
		v := it.b0
		it.b0, it.b1 = it.b.Next()
		return v, true
	}
	return 0, false
}
