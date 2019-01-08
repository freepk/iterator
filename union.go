package iterator

type UnionIterator struct {
	a []Iterator
	v []int
	n int
}

func NewUnionIterator(a []Iterator) *UnionIterator {
	n := len(a)
	v := make([]int, n)
	it := &UnionIterator{a: a, v: v}
	it.n = first(it.a, it.v)
	return it
}

func (it *UnionIterator) Reset() {
	n := len(it.a)
	for i := 0; i < n; i++ {
		it.a[i].Reset()
	}
	it.n = first(it.a, it.v)
}

func (it *UnionIterator) Next() (int, bool) {
	if it.n == 0 {
		return 0, false
	}
	v := it.v[0]
	i := 1
	for i < it.n {
		if v > it.v[i] {
			v = it.v[i]
		}
		i++
	}
	i = 0
	for i < it.n {
		if it.v[i] > v {
			i++
			continue
		}
		x, ok := it.a[i].Next()
		if !ok {
			it.n--
			if i != it.n {
				it.a[i], it.a[it.n] = it.a[it.n], it.a[i]
				it.v[i], it.v[it.n] = it.v[it.n], it.v[i]
			}
			continue
		}
		if x == v {
			continue
		}
		it.v[i] = x
		i++
	}
	return v, true
}

type ArrUnionIterator struct {
	a [][]int
	v []int
	i []int
	n int
}

func NewArrUnionIterator(a [][]int) *ArrUnionIterator {
	n := len(a)
	v := make([]int, n)
	i := make([]int, n)
	it := &ArrUnionIterator{a: a, v: v, i: i}
	it.n = firsta(it.a, it.v, it.i)
	return it
}

func (it *ArrUnionIterator) Reset() {
	n := len(it.a)
	for i := 0; i < n; i++ {
		it.i[i] = 0
	}
	it.n = firsta(it.a, it.v, it.i)
}

func (it *ArrUnionIterator) Next() (int, bool) {
	if it.n == 0 {
		return 0, false
	}
	v := it.v[0]
	i := 1
	for i < it.n {
		if v > it.v[i] {
			v = it.v[i]
		}
		i++
	}
	i = 0
	for i < it.n {
		if it.v[i] > v {
			i++
			continue
		}
		k := it.i[i]
		if k >= len(it.a[i]) {
			it.n--
			if i != it.n {
				it.a[i], it.a[it.n] = it.a[it.n], it.a[i]
				it.v[i], it.v[it.n] = it.v[it.n], it.v[i]
				it.i[i], it.i[it.n] = it.i[it.n], it.i[i]
			}
			continue
		}
		it.i[i]++
		x := it.a[i][k]
		if x == v {
			continue
		}
		it.v[i] = x
		i++
	}
	return v, true
}
