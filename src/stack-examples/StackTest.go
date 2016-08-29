package main

import (
	"fmt"
	"stack"
)

func main() {

	// get the  Integer Stack

	arrayStack := stack.GetnewOne(5)
	// Testing Size
	fmt.Printf("Stack Size %d : \n ", arrayStack.Size())
	// using of POP opeerations
	for i := 1; i <= 5; i++ {
		anyError := arrayStack.Push(i * 25)
		fmt.Printf("Element pushed: %d \n  ", i*25)
		if anyError != nil {
			fmt.Println(anyError)
		}
	}
	// Testing Size O(1)
	fmt.Printf("Stack Size %d :  \n", arrayStack.Size())
	// Usage of POP operation
	for i := 0; i < 5; i++ {
		popElement, anyError := arrayStack.Pop()
		if anyError == nil {
			fmt.Printf("Element Popped : %d \n  ", popElement)
		} else {
			fmt.Println(anyError)
		}

	}

	// get the  Integer Stack

	stringStack := stack.GetnewOne(5)
	// Testing Size
	fmt.Printf("Stack Size %d : \n ", stringStack.Size())
	// using of POP opeerations
	var words []string = []string{"A", "B", "C", "D", "F"}

	for _, word := range words {

		anyError := stringStack.Push(word)
		fmt.Printf("Element pushed : %s \n  ", word)
		if anyError != nil {
			fmt.Println(anyError)
		}
	}
	// Testing Size O(1)
	fmt.Printf("Stack Size %d :  \n", stringStack.Size())
	// Usage of POP operation
	for i := 0; i < 5; i++ {
		value, anyError := stringStack.Pop()
		if anyError == nil {
			fmt.Printf("Element Popped : %s \n  ", value)
		} else {
			fmt.Println(anyError)
		}
	}
	fmt.Printf("Stack Size : %d  \n ", stringStack.Size())
}
