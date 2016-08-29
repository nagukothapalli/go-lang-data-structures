package main

import (
	"fmt"
)

//Main
func main() {

	intr := " ## Single Linked List ## \n"

	fmt.Println(intr)
	// get The new Single List
	list := IList(GetNewList())
	// Adding elememts
	list.AddLast("Kothapalli")
	list.AddFirst("Nagu")
	list.AddFirst("Lives in Pune")
	list.AddLast("Works at card") // add First
	list.PrintList()
	// Print the Size
	fmt.Printf("Linked List Size : %d \n", list.Size())
	//	list.removeFirst()
	// Print the Size
	fmt.Printf("Linked List Size : %d \n", list.Size())
	// Print the Size
	fmt.Printf(" Truth Valeu: %t \n", list.IsEmpty())

	fmt.Println()
	// After Removel of Some thing
	list.RemoveFirst()

	list.PrintList()
	// Print the Size
	fmt.Printf("Linked List Size : %d \n", list.Size())
	//	list.removeFirst()
	// Print the Size
	fmt.Printf("Linked List Size : %d \n", list.Size())
	// Print the Size
	fmt.Printf(" Truth Valeu: %t \n", list.IsEmpty())

}
