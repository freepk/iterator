package iterator

type Iterator interface {
	Reset()
	ResetToEnd()
	Next() (int, bool)
	Prev() (int, bool)
}
