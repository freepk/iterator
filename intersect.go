package iterator

type IntersectIterator struct {
	a []Iterator
}

func NewIntersectIterator(a []Iterator) *IntersectIterator {
	return &IntersectIterator{a: a}
}

func (it *IntersectIterator) Reset() {
	n := len(it.a)
	for i := 0; i < n; i++ {
		it.a[i].Reset()
	}
}

func (it *IntersectIterator) Next() (int, bool) {
	n := len(it.a)
	v, ok := it.a[0].Next()
	if !ok {
		return 0, false
	}
	i := 1
	for i < n {
		x, ok := it.a[i].Next()
		if !ok {
			return 0, false
		}
		if x < v {
			continue
		}
		if x > v {
			it.a[0], it.a[i] = it.a[i], it.a[0]
			i = 1
			v = x
			continue
		}
		i++
	}
	return v, true
}

type ArrIntersectIterator struct {
	a [][]int
	i []int
}

func NewArrIntersectIterator(a [][]int) *ArrIntersectIterator {
	n := len(a)
	i := make([]int, n)
	return &ArrIntersectIterator{a: a, i: i}
}

func (it *ArrIntersectIterator) Reset() {
	n := len(it.i)
	for i := 0; i < n; i++ {
		it.i[i] = 0
	}
}

func (it *ArrIntersectIterator) Next() (int, bool) {
	n := len(it.a)
	k := it.i[0]
	if k >= len(it.a[0]) {
		return 0, false
	}
	it.i[0]++
	v := it.a[0][k]
	i := 1
	for i < n {
		k = it.i[i]
		if k >= len(it.a[i]) {
			return 0, false
		}
		it.i[i]++
		x := it.a[i][k]
		if x < v {
			continue
		}
		if x > v {
			it.a[0], it.a[i] = it.a[i], it.a[0]
			it.i[0], it.i[i] = it.i[i], it.i[0]
			i = 1
			v = x
			continue
		}
		i++
	}
	return v, true
}
