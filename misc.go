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

func firsta(a [][]int, v []int, k []int) int {
	n := len(a)
	i := 0
	for i < n {
		if len(a[i]) > 0 {
			v[i] = a[i][0]
			k[i] = 1
			i++
		} else {
			n--
			a[i] = a[n]
		}
	}
	return n
}

func swap(a []Iterator, v []int, i, j int) {
	if i != j {
		a[i], a[j] = a[j], a[i]
		v[i], v[j] = v[j], v[i]
	}
}

func swapa(a [][]int, v []int, k []int, i, j int) {
	if i != j {
		a[i], a[j] = a[j], a[i]
		v[i], v[j] = v[j], v[i]
		k[i], k[j] = k[j], k[i]
	}
}
