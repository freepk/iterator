package iterator

type UnionIterator struct {
	a []Iterator
	v []int
	p int
}

func NewUnionIterator(a []Iterator) *UnionIterator {
	it := &UnionIterator{a: a}
	it.first()
	return it
}

func (it *UnionIterator) first() {
	it.v = it.v[:0]
	n := len(it.a)
	m := 0
	i := 0
	for i < n {
		v, ok := it.a[i].Next()
		if ok {
			j := 0
			for j < m {
				if v < it.v[j] {
					break
				}
				j++
			}
			m++
			it.v = append(it.v, 0)
			copy(it.v[j+1:], it.v[j:])
			it.v[j] = v
			it.a[i], it.a[j] = it.a[j], it.a[i]
			i++
		} else {
			n--
			it.a[i] = it.a[n]
		}
	}
	it.a = it.a[:n]
	it.p = 0
}

func (it *UnionIterator) Reset() {
	n := len(it.a)
	for i := 0; i < n; i++ {
		it.a[i].Reset()
	}
	it.first()
}

func (it *UnionIterator) Next() (int, bool) {
	n := len(it.a)
	p := it.p
	if p >= n {
		return 0, false
	}
	v := it.v[p]
	x, ok := it.a[p].Next()
	if ok {
		it.v[p] = x
		p++
		for p < n {
			if x < it.v[p] {
				break
			}
			it.a[p-1], it.a[p] = it.a[p], it.a[p-1]
			it.v[p-1], it.v[p] = it.v[p], it.v[p-1]
			p++
		}
	} else {
		it.p++
	}
	return v, true
}
