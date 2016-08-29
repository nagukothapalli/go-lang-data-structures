package stack

// Allowed operations using stack
type IStack interface {
	Push(element interface{}) error
	Pop() (interface{}, error)
	IsEmpty() bool
	Size() int
	Peek() (interface{}, error)
}
