package iterator

type UnionIterator struct {
	a []Iterator
	v []int
	n int
}

func NewUnionIterator(a []Iterator) *UnionIterator {
	it := &UnionIterator{a: a}
	it.first()
	return it
}

func (it *UnionIterator) first() {
	it.v = it.v[:0]
	n := len(it.a)
	i := 0
	for i < n {
		v, ok := it.a[i].Next()
		if !ok {
			n--
			it.a[i] = it.a[n]
			continue
		}
		it.v = append(it.v, v)
		i++
	}
	it.n = n
}

func (it *UnionIterator) Reset() {
	n := len(it.a)
	for i := 0; i < n; i++ {
		it.a[i].Reset()
	}
	it.first()
}

func (it *UnionIterator) Next() (int, bool) {
	if it.n == 0 {
		return 0, false
	}
	i := 0
	v := it.v[0]
	for j := 1; j < it.n; j++ {
		if v > it.v[j] {
			i = j
			v = it.v[j]
		}
	}
	x, ok := it.a[i].Next()
	if ok {
		it.v[i] = x
	} else {
		it.n--
		if i != it.n {
			it.a[i], it.a[it.n] = it.a[it.n], it.a[i]
			it.v[i], it.v[it.n] = it.v[it.n], it.v[i]
		}
	}
	return v, true
}
