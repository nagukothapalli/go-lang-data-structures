package main

// Allowed operations using stack
type Stack interface {
	Push(node *Node) error
	Pop() (*Node, error)
	IsEmpty() bool
	Size() int
	Peek() (*Node, error)
	GetnewOne(capacity uint) *Stack
}
