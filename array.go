package iterator

type ArrayIter struct {
	a []int
	n int
	i int
}

func NewArrayIter(a []int) *ArrayIter {
	n := len(a)
	return &ArrayIter{a: a, n: n, i: 0}
}

func (it *ArrayIter) Reset() {
	it.i = 0
}

func (it *ArrayIter) Next() (int, bool) {
	i := it.i
	if i < it.n {
		it.i++
		return it.a[i], true
	}
	return 0, false
}
