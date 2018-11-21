package main

type Aggregate interface {
	iterator() Iterator
}
