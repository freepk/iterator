package iterator

func first(a []Iterator, v []int) int {
	n := len(a)
	i := 0
	for i < n {
		x, ok := a[i].Next()
		if ok {
			v[i] = x
			i++
		} else {
			n--
			a[i] = a[n]
		}
	}
	return n
}

func firsta(a [][]int, v []int, i []int) int {
	n := len(a)
	j := 0
	for j < n {
		if len(a[j]) > 0 {
			v[j] = a[j][0]
			i[j] = 1
			j++
		} else {
			n--
			a[j] = a[n]
		}
	}
	return n
}
