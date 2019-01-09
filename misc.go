package iterator

import (
	"math/rand"
	"sort"
	"time"
)

func init() {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)
}

func randArray(n int) []int {
	if n == 0 {
		return []int{}
	}
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = rand.Intn(n * 2)
	}
	sort.Ints(r)
	return r
}

func isEqual(a, b []int) bool {
	n := len(a)
	if n != len(b) {
		return false
	}
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
