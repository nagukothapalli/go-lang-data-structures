package queue

//  Operations on queue
type IQueue interface {
	
	EnQueue(element interface{}) error
	DeQueue() (interface{},error) 
	Front() (interface{},error)
	Rear() (interface{},error)
	Size() int
	IsEmpty() bool
	
}
