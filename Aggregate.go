package main

// Aggregate interface
type Aggregate interface {
	Iterator() Iterator
}
