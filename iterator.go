package iterator

import (
	"errors"
)

var (
	iterInvalidParamErr = errors.New("Invalid parameter")
)

type Iterator interface {
	Reset()
	Next() (int, bool)
}

type unionIter struct {
}

func newUnionIter(a Iterator, b Iterator) Iterator {
	return nil
}

type unionAllIter struct {
}

func newUnionAllIter(a Iterator, b Iterator) Iterator {
	return nil
}

type exceptIter struct {
}

func newExceptIter(a Iterator, b Iterator) Iterator {
	return nil
}
