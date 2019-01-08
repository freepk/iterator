package iterator

type ExceptIterator struct {
	a Iterator
	b []Iterator
	v []int
	n int
}

func NewExceptIterator(a Iterator, b []Iterator) *ExceptIterator {
	n := len(b)
	v := make([]int, n)
	it := &ExceptIterator{a: a, b: b, v: v}
	it.n = first(it.b, it.v)
	return it
}

func (it *ExceptIterator) Reset() {
	it.a.Reset()
	n := len(it.b)
	for i := 0; i < n; i++ {
		it.b[i].Reset()
	}
	it.n = first(it.b, it.v)
}

func (it *ExceptIterator) Next() (int, bool) {
	v, ok := it.a.Next()
	if !ok {
		return 0, false
	}
	i := 0
	for i < it.n {
		if v == it.v[i] {
			v, ok = it.a.Next()
			if !ok {
				return 0, false
			}
			i = 0
			continue
		}
		if v > it.v[i] {
			x, ok := it.b[i].Next()
			if !ok {
				it.n--
				swap(it.b, it.v, i, it.n)
				continue
			}
			it.v[i] = x
			continue
		}
		i++
	}
	return v, true
}
