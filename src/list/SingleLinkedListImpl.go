package main

import "errors"

type SingleLinkedList struct {
	head *SingleLinkedListNode
	tail *SingleLinkedListNode
	size int
}

//Implementaions
// Get the new Lit
func getNewList() *SingleLinkedList {
	return &SingleLinkedList{head: nil, tail: nil, size: 0}
}

// get the size
func (list *SingleLinkedList) Size() int {
	return list.size
}

// check wether List is empty or not

func (list *SingleLinkedList) isEmpty() bool {
	return (list.head == nil && list.tail == nil)
}

// add the  new element as a first node

func (list *SingleLinkedList) addFirst(node *SingleLinkedListNode) {

	if list.isEmpty() {
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
func (list *SingleLinkedList) addLast(node *SingleLinkedListNode) {

	if list.isEmpty() {
		list.head = node
		list.tail = node

	} else {
		list.tail.next = node
		list.tail = node

	}
	list.size++

}

// remove the first  element

func (list *SingleLinkedList) removeFirst() error {

	if list.isEmpty() {
		return errors.New("No elements in list to delete!!")
	}
	futurHead := list.head.next
	list.head = nil // this bacaues for flag as  garbage
	list.head = futurHead
	decrementSize(list)
	return nil

}

func decrementSize(list *SingleLinkedList) {
	list.size--
	if list.size == 0 {
		list.head = nil
		list.tail = nil
	}

}
