package iterator

type UnionAllIterator struct {
	a []Iterator
	v []int
	n int
}

func NewUnionAllIterator(a []Iterator) *UnionAllIterator {
	n := len(a)
	v := make([]int, n)
	it := &UnionAllIterator{a: a, v: v}
	it.n = first(it.a, it.v)
	return it
}

func (it *UnionAllIterator) Reset() {
	n := len(it.a)
	for i := 0; i < n; i++ {
		it.a[i].Reset()
	}
	it.n = first(it.a, it.v)
}

func (it *UnionAllIterator) Next() (int, bool) {
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

type ArrUnionAllIterator struct {
	a [][]int
	v []int
	i []int
	n int
}

func NewArrUnionAllIterator(a [][]int) *ArrUnionAllIterator {
	n := len(a)
	v := make([]int, n)
	i := make([]int, n)
	it := &ArrUnionAllIterator{a: a, v: v, i: i}
	it.n = firsta(it.a, it.v, it.i)
	return it
}

func (it *ArrUnionAllIterator) Reset() {
	n := len(it.a)
	for i := 0; i < n; i++ {
		it.i[i] = 0
	}
	it.n = firsta(it.a, it.v, it.i)
}

func (it *ArrUnionAllIterator) Next() (int, bool) {
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
	k := it.i[i]
	if k < len(it.a[i]) {
		it.v[i] = it.a[i][k]
		it.i[i]++
	} else {
		it.n--
		if i != it.n {
			it.a[i], it.a[it.n] = it.a[it.n], it.a[i]
			it.v[i], it.v[it.n] = it.v[it.n], it.v[i]
			it.i[i], it.i[it.n] = it.i[it.n], it.i[i]
		}
	}
	return v, true
}
