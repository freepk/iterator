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

func (it *ArrayIter) ResetToEnd() {
	it.i = it.n + 1
}

func (it *ArrayIter) Next() (int, bool) {
	i := it.i
	it.i++
	if i < it.n {
		return it.a[i], true
	}
	return 0, false
}

func (it *ArrayIter) Prev() (int, bool) {
	it.i--

	if it.i - 1 < 0 {
		return 0, false
	}

	return it.a[it.i - 1], true
}