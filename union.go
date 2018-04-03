package iterator

type UnionIterator struct {
	initIterators []Iterator
	initValues    []int
	iterators     []Iterator
	values        []int
}

func NewUnionIterator(iterators []Iterator) *UnionIterator {
	var initIterators []Iterator
	var initValues []int
	size := len(iterators)
	for i := 0; i < size; i++ {
		iterators[i].Reset()
		value, ok := iterators[i].Next()
		if ok {
			initIterators = append(initIterators, iterators[i])
			initValues = append(initValues, value)
		}
	}
	return &UnionIterator{
		initIterators: initIterators,
		initValues:    initValues,
		iterators:     initIterators,
		values:        initValues,
	}
}

func (it *UnionIterator) Reset() {
	size := len(it.initIterators)
	for i := 0; i < size; i++ {
		it.initIterators[i].Reset()
		it.initValues[i], _ = it.initIterators[i].Next()
	}
	it.iterators = it.initIterators
	it.values = it.initValues
}

func (it *UnionIterator) Next() (int, bool) {
	size := len(it.iterators)
	if size == 0 {
		return 0, false
	}
	offset := 0
	advice := it.values[0]
	for i := 1; i < size; i++ {
		if advice > it.values[i] {
			offset = i
			advice = it.values[i]
		}
	}
	value, ok := it.iterators[offset].Next()
	if ok {
		it.values[offset] = value
	} else {
		it.iterators = append(it.iterators[:offset], it.iterators[offset+1:]...)
		it.values = append(it.values[:offset], it.values[offset+1:]...)
	}
	return advice, true
}
