package main

import (
	"fmt"
)

//  --  Main
func main() {

	intr := `########################
     ## Single Linked List ##
     
     ######################### `

	fmt.Println(intr)
	// get The new Single List
	list := getNewList()

	//	list.addFirst(&SingleLinkedListNode{element: "Nagu", next: nil})
	list.addLast(&SingleLinkedListNode{element: "Kothapalli", next: nil})
	list.addFirst(&SingleLinkedListNode{element: "Nagu", next: nil})
	list.addFirst(&SingleLinkedListNode{element: "Lives in Pune", next: nil})
	list.addLast(&SingleLinkedListNode{element: "Works at card", next: nil})

	// add First
	printList(list)
	// Print the Size
	fmt.Printf("Linked List Size : %d \n", list.Size())
	//	list.removeFirst()
	// Print the Size
	fmt.Printf("Linked List Size : %d \n", list.Size())
	// Print the Size
	fmt.Printf(" Truth Valeu: %t \n", list.isEmpty())

	fmt.Println("\n\n\n")
	// After Removel of Some thing
	list.removeFirst()

	printList(list)
	// Print the Size
	fmt.Printf("Linked List Size : %d \n", list.Size())
	//	list.removeFirst()
	// Print the Size
	fmt.Printf("Linked List Size : %d \n", list.Size())
	// Print the Size
	fmt.Printf(" Truth Valeu: %t \n", list.isEmpty())

}

func printList(list *SingleLinkedList) {

	if list.isEmpty() {
		fmt.Println(" No elements to print")
	} else {
		listCrawler := list.head
		for listCrawler != nil {
			// Asuming All the Elemens are in list are strings
			fmt.Printf("Element: %s \n", listCrawler.element)
			listCrawler = listCrawler.next
		}
	}
}
