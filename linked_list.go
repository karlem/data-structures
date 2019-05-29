package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() LinkedList {
	return LinkedList{}
}

func (li *LinkedList) addAtFront(data interface{}) {
	if li.head == nil {
		li.head = &Node{data, nil}
		return
	}

	newNode := &Node{data, li.head}
	li.head = newNode
}

func (li *LinkedList) addAtEnd(data interface{}) {
	if li.head == nil {
		li.head = &Node{data, nil}
		return
	}

	node := li.head
	for {
		if node.next == nil {
			node.next = &Node{data, nil}
			break
		}
		node = node.next
	}
}

func (li *LinkedList) print() {
	if li.head == nil {
		fmt.Println("List is empty")
	}

	node := li.head
	for node != nil {
		fmt.Printf("=> %v ", node.data)
		node = node.next
	}
}
