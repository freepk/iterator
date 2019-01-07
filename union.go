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
	it.first()
	return it
}

func (it *UnionIterator) first() {
	n := len(it.a)
	i := 0
	for i < n {
		v, ok := it.a[i].Next()
		if ok {
			it.v[i] = v
			i++
		} else {
			n--
			it.a[i] = it.a[n]
		}
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
	it.first()
	return it
}

func (it *ArrUnionIterator) first() {
	n := len(it.a)
	i := 0
	for i < n {
		if len(it.a[i]) > 0 {
			it.v[i] = it.a[i][0]
			it.i[i] = 1
			i++
		} else {
			n--
			it.a[i] = it.a[n]
		}
	}
	it.n = n
}

func (it *ArrUnionIterator) Reset() {
	n := len(it.a)
	for i := 0; i < n; i++ {
		it.i[i] = 0
	}
	it.first()
}

func (it *ArrUnionIterator) Next() (int, bool) {
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
