package iterator

type IntersectIterator struct {
	iterators []Iterator
}

func NewIntersectIterator(iterators []Iterator) *IntersectIterator {
	return &IntersectIterator{iterators: iterators}
}

func (it *IntersectIterator) Reset() {
	size := len(it.iterators)
	for i := 0; i < size; i++ {
		it.iterators[i].Reset()
	}
}

func (it *IntersectIterator) Next() (int, bool) {
	size := len(it.iterators)
	advice, ok := it.iterators[0].Next()
	if !ok {
		return 0, false
	}
	i := 1
	equals := 1
	for equals != size {
		value, ok := it.iterators[i].Next()
		if !ok {
			return 0, false
		}
		for advice > value {
			value, ok = it.iterators[i].Next()
			if !ok {
				return 0, false
			}
		}
		if advice < value {
			advice = value
			i = 0
			equals = 1
			continue
		}
		equals++
		i++
	}
	return advice, true
}
