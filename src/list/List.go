package main

type List interface {
	Size() int
	addFirtst()
	addlast()
	isEmpty() bool
	removeFirst() error
	getNewList() *SingleLinkedList
}
