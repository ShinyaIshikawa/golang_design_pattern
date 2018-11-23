package main

// Iterator interface
type Iterator interface {
	hasNext() bool
	next() Aggregate
}
