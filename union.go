package iterator

type UnionIter struct {
	a  Iterator
	b  Iterator
	a0 int
	a1 bool
	b0 int
	b1 bool
}

func NewUnionIter(a, b Iterator) *UnionIter {
	it := &UnionIter{a: a, b: b}
	it.a0, it.a1 = it.a.Next()
	it.b0, it.b1 = it.b.Next()
	return it
}

func (it *UnionIter) Reset() {
	it.a.Reset()
	it.b.Reset()
	it.a0, it.a1 = it.a.Next()
	it.b0, it.b1 = it.b.Next()
}

func (it *UnionIter) Next() (int, bool) {
	v := 0
	if it.a1 && it.b1 {
		v = it.a0
		if v > it.b0 {
			v = it.b0
		}
	} else if it.a1 {
		v = it.a0
	} else if it.b1 {
		v = it.b0
	} else {
		return 0, false
	}
	for it.a0 == v && it.a1 {
		it.a0, it.a1 = it.a.Next()
	}
	for it.b0 == v && it.b1 {
		it.b0, it.b1 = it.b.Next()
	}
	return v, true
}
