package main

type IList interface {
	Size() int
	AddFirst(element interface{})
	AddLast(element interface{})
	IsEmpty() bool
	RemoveFirst() error
	PrintList()
	//GetHeadElement() interface{}
	//GetTailElement() interface{}
}
