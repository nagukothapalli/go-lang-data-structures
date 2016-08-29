package queue

import "errors"
// Define the Type of Queue
type ArrayQueue struct {
    front int
    rear int
    size uint
    capacity uint
    elements *[] interface{}
}

var (
   _queueOverFlow := errors.New("Queue overFlow ( Queue got full) !!")
   _queueUnderFlow := errors.New("Queue Empty !!")
)
//Get the new Queue
func GetNewQueue(capacity uint) *ArrayQueue{
return &ArrayQueue
{
    front:-1,
    rear:-1,
    size:0,
    capacity:capacity,elements:make(*[] interface{},capacity)
}
}
// Get the Size
func (queue *ArrayQueue) Size() uint {
    return queue.size
}
// IsEmpty()
func (queue *ArrayQueue) IsEmpty() bool {
    return queue.size == 0 
}

//Enqueue()
func (queue *ArrayQueue) EnQueue(element interface{}) error {
     if queue.size == queue.capacity {
        return _queueOverFlow
    } else if queue.IsEmpty() {
        queue.front = 0
        queue.rear  = 0
    } else {
        queue.rear ++
    }
queue.elements[queue.rear] = element
queue.size ++
return nil
}

func (queue *ArrayQueue) DeQueue() (element interface{}, error) {
   element interface{}
    if queue.IsEmpty() {
        return (nil,_queueUnderFlow)
    }else if queue.front == queue.rear {
        element = queue.elements[queue.front]
        queue.elements[queue.front]=nil //mark as the garbage
        queue.front-- 
        queue.rear--
        

    } else {
element = queue.elements[queue.front]
queue.elements[queue.front]=nil
        queue.front ++

    }
return element,nil
}
// To get the Front element
func (queue *ArrayQueue) Front() (interface{},error) {
    if queue.IsEmpty() {
    return (nil,_queueUnderFlow)
    }

return (queue.elements[queue.front],nil)

}


func (queue *ArrayQueue) Rear() (interface{},error) {
    if queue.IsEmpty() {
    return (nil,_queueUnderFlow)
    }
return (queue.elements[queue.rear],nil)

}