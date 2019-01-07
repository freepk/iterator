package iterator

type IntersectIterator struct {
	arr []Iterator
}

func NewIntersectIterator(arr []Iterator) *IntersectIterator {
	return &IntersectIterator{arr: arr}
}

func (it *IntersectIterator) Reset() {
	n := len(it.arr)
	for i := 0; i < n; i++ {
		it.arr[i].Reset()
	}
}

func (it *IntersectIterator) Next() (int, bool) {
	n := len(it.arr)
	v, ok := it.arr[0].Next()
	if !ok {
		return 0, false
	}
	i := 1
	for i < n {
		x, ok := it.arr[i].Next()
		if !ok {
			return 0, false
		}
		if x < v {
			continue
		}
		if x > v {
			it.arr[0], it.arr[i] = it.arr[i], it.arr[0]
			i = 1
			v = x
			continue
		}
		i++
	}
	return v, true
}

type ArrIntersectIterator struct {
	arr [][]int
	i   []int
}

func NewArrIntersectIterator(arr [][]int) *ArrIntersectIterator {
	n := len(arr)
	i := make([]int, n)
	return &ArrIntersectIterator{arr: arr, i: i}
}

func (it *ArrIntersectIterator) Reset() {
	n := len(it.i)
	for i := 0; i < n; i++ {
		it.i[i] = 0
	}
}

func (it *ArrIntersectIterator) Next() (int, bool) {
	n := len(it.arr)
	k := it.i[0]
	if k >= len(it.arr[0]) {
		return 0, false
	}
	it.i[0]++
	v := it.arr[0][k]
	i := 1
	for i < n {
		k = it.i[i]
		if k >= len(it.arr[i]) {
			return 0, false
		}
		it.i[i]++
		x := it.arr[i][k]
		if x < v {
			continue
		}
		if x > v {
			it.arr[0], it.arr[i] = it.arr[i], it.arr[0]
			it.i[0], it.i[i] = it.i[i], it.i[0]
			i = 1
			v = x
			continue
		}
		i++
	}
	return v, true
}
