package main

import (
	"errors"
	"fmt"
)

type SingleLinkedList struct {
	head *SingleLinkedListNode
	tail *SingleLinkedListNode
	size int
}

//Implementaions
// Get the new Lit
func GetNewList() *SingleLinkedList {
	return &SingleLinkedList{head: nil, tail: nil, size: 0}
}

// get the size
func (list *SingleLinkedList) Size() int {
	return list.size
}

// check wether List is empty or not

func (list *SingleLinkedList) IsEmpty() bool {
	return (list.head == nil && list.tail == nil)
}

// add the  new element as a first node

func (list *SingleLinkedList) AddFirst(element interface{}) {
	node := InitNode(element)
	if list.IsEmpty() {
		list.head = node
		list.tail = node

	} else {
		featurhead := list.head
		list.head = node
		list.head.next = featurhead
	}
	list.size++
}

// add the new element as a last node
func (list *SingleLinkedList) AddLast(element interface{}) {
	node := InitNode(element)
	if list.IsEmpty() {
		list.head = node
		list.tail = node

	} else {
		list.tail.next = node
		list.tail = node

	}
	list.size++

}

// remove the first  element

func (list *SingleLinkedList) RemoveFirst() error {

	if list.IsEmpty() {
		return errors.New("No elements in list to delete!!")
	}
	futurHead := list.head.next
	list.head = nil // this bacaues for flag as  garbage
	list.head = futurHead
	DecrementSize(list)
	return nil

}
func (list *SingleLinkedList) PrintList() {

	if list.IsEmpty() {
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

// Some utils functions
func InitNode(element interface{}) *SingleLinkedListNode {
	return &SingleLinkedListNode{element: element, next: nil}
}

func DecrementSize(list *SingleLinkedList) {
	list.size--
	if list.size == 0 {
		list.head = nil
		list.tail = nil
	}
}
