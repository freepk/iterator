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
		swap(it.a, it.v, i, it.n)
	}
	return v, true
}

type ArrUnionAllIterator struct {
	a [][]int
	v []int
	k []int
	n int
}

func NewArrUnionAllIterator(a [][]int) *ArrUnionAllIterator {
	n := len(a)
	v := make([]int, n)
	k := make([]int, n)
	it := &ArrUnionAllIterator{a: a, v: v, k: k}
	it.n = firsta(it.a, it.v, it.k)
	return it
}

func (it *ArrUnionAllIterator) Reset() {
	n := len(it.a)
	for i := 0; i < n; i++ {
		it.k[i] = 0
	}
	it.n = firsta(it.a, it.v, it.k)
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
	k := it.k[i]
	if k < len(it.a[i]) {
		it.v[i] = it.a[i][k]
		it.k[i]++
	} else {
		it.n--
		if i != it.n {
			swapa(it.a, it.v, it.k, i, it.n)
		}
	}
	return v, true
}
