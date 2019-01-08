package iterator

import (
	"testing"

	"github.com/freepk/arrays"
)

func TestExcept(t *testing.T) {
	a := []int{0, 100, 100, 200, 300, 350, 400}
	b := []int{400, 500}
	c := []int{200, 400}
	d := combineArrays([][]int{a, b, c}, arrays.Except)
	if !arrays.IsEqual([]int{0, 100, 100, 300, 350}, d) {
		t.Fail()
	}
}
