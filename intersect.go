package iterator

type IntersectIterator struct {
	iterators []Iterator
}

func NewIntersectIterator(iterators []Iterator) *IntersectIterator {
	iter := &IntersectIterator{iterators: iterators}
	iter.Reset()
	return iter
}

func (it *IntersectIterator) Reset() {
	size := len(it.iterators)
	for i := 0; i < size; i++ {
		it.iterators[i].Reset()
	}
}

func (it *IntersectIterator) Next() (int, bool) {
	size := len(it.iterators)
	i := 0
	advice, ok := it.iterators[i].Next()
	if !ok {
		return 0, false
	}
	skip := i
	for i < size {
		if i == skip {
			i++
			continue
		}
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
			skip = i
			i = 0
			continue
		}
		i++
	}
	return advice, true
}
