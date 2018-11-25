package iterator

// Iterator interface
type Iterator interface {
	hasNext() bool
	next() Aggregate
}
