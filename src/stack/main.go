package main

import (
	"fmt"
)

func main() {

	// get the  Integer Stack

	stack := GetnewOne(5)
	// Testing Size
	fmt.Printf("Stack Size %d : \n ", stack.Size())
	// using of POP opeerations
	for i := 1; i <= 5; i++ {
		anyError := stack.Push(&Node{element: i * 25})
		fmt.Printf("Element pushed: %d \n  ", i*25)
		if anyError != nil {
			fmt.Println(anyError)
		}
	}
	// Testing Size O(1)
	fmt.Printf("Stack Size %d :  \n", stack.Size())
	// Usage of POP operation
	for i := 0; i < 5; i++ {
		value, anyError := stack.Pop()
		if anyError == nil {
			fmt.Printf("Element Popped : %d \n  ", value.element)
		} else {
			fmt.Println(anyError)
		}

	}

	// get the  Integer Stack

	stringStack := GetnewOne(5)
	// Testing Size
	fmt.Printf("Stack Size %d : \n ", stringStack.Size())
	// using of POP opeerations
	var words []string = []string{"A", "B", "C", "D", "F"}

	for _, word := range words {

		anyError := stringStack.Push(&Node{element: word})
		fmt.Printf("Element pushed : %s \n  ", word)
		if anyError != nil {
			fmt.Println(anyError)
		}
	}
	// Testing Size O(1)
	fmt.Printf("Stack Size %d :  \n", stack.Size())
	// Usage of POP operation
	for i := 0; i < 5; i++ {
		value, anyError := stringStack.Pop()
		if anyError == nil {
			fmt.Printf("Element Popped : %s \n  ", value.element)
		} else {
			fmt.Println(anyError)
		}
	}
	fmt.Printf("Stack Size : %d  \n ", stringStack.Size())
}
