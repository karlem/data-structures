package main

import (
	"errors"
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	length int
	head   *Node
}

func NewLinkedList() LinkedList {
	return LinkedList{}
}

func (li *LinkedList) AddAtFront(data interface{}) {
	var newNode *Node
	if li.head == nil {
		newNode = &Node{data, nil}
	} else {
		newNode = &Node{data, li.head}
	}

	li.head = newNode
	li.length++
}

func (li *LinkedList) AddAtEnd(data interface{}) {
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

	li.length++
}

func (li *LinkedList) Delete(key interface{}) {
	node := li.head
	var prev *Node

	for node != nil {
		if node.data == key {
			if prev == nil {
				li.head = node.next
			} else {
				prev.next = node.next
			}
			li.length--
		}

		prev = node
		node = node.next
	}
}

func (li *LinkedList) Last() (interface{}, error) {
	if li.head == nil {
		return nil, errors.New("List is empty")
	}

	node := li.head
	for node.next != nil {
		node = node.next
	}

	return node.data, nil
}

func (li *LinkedList) First() (interface{}, error) {
	if li.head == nil {
		return nil, errors.New("List is empty")
	}

	node := li.head
	for node.next != nil {
		node = node.next
	}

	return node.data, nil
}

func (li *LinkedList) Length() int {
	return li.length
}

func (li *LinkedList) ToArray() ([]interface{}, error) {
	if li.head == nil {
		return nil, errors.New("List is empty")
	}

	outArray := []interface{}{}
	node := li.head

	for node != nil {
		outArray = append(outArray, node.data)
		node = node.next
	}
	return outArray, nil
}

func (li *LinkedList) Print() {
	if li.head == nil {
		fmt.Println("List is empty")
	}

	node := li.head
	for node != nil {
		fmt.Printf("=> %v ", node.data)
		node = node.next
	}
	fmt.Println("")
}
