package stack

import "errors"

//Defining the Array  based stack
type ArrayStack struct {
	nodes    []*interface{}
	top      int
	capacity int
}

// Defining the  possible  errors may occur .

var (
	_underflowEror = errors.New("Stack UnderFlow Error !!")
	_overflowEror  = errors.New("Stack overFlow Error !!")
)

// Define the Stack operations
func GetnewOne(capacity int) *ArrayStack {
	newStack := ArrayStack{nodes: make([]*interface{}, capacity), capacity: capacity, top: -1}
	return &newStack
}

// check wether stack is empty or not
func (stack *ArrayStack) IsEmpty() bool {
	return (stack.top == -1)
}

// Push the  elements into stack
func (stack *ArrayStack) Push(element interface{}) error {

	dummyTop := stack.top
	dummyTop = dummyTop + 1

	if dummyTop >= stack.capacity {
		return _overflowEror
	}
	stack.top = dummyTop
	stack.nodes[stack.top] = &element
	return nil
}

// POP the elements from stack

func (stack *ArrayStack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, _underflowEror
	}
	topNode := stack.nodes[stack.top]
	stack.top--
	return *topNode, nil

}

// Implemeting The Size
func (stack *ArrayStack) Size() int {
	if stack.IsEmpty() {
		return 0
	}
	return stack.top + 1
}

// Implemeting the peek Method
func (stack *ArrayStack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, _underflowEror
	}
	topNode := stack.nodes[stack.top]
	return *topNode, nil
}
