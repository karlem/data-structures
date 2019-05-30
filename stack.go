package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	top *Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (st *Stack) IsEpmty() bool {
	return st.top == nil
}

func (st *Stack) Push(data interface{}) {
	var newNode *Node

	if st.top == nil {
		newNode = &Node{data, nil}
	} else {
		newNode = &Node{data, st.top}
	}

	st.top = newNode
}

func (st *Stack) Pop() (interface{}, error) {
	if st.top == nil {
		return nil, errors.New("Stack is empty")
	}

	out := st.top
	st.top = out.next

	return out.data, nil
}

func (st *Stack) Print() {
	if st.top == nil {
		fmt.Println("Stack is empty")
	}

	node := st.top
	for node != nil {
		fmt.Printf("=> %v ", node.data)
		node = node.next
	}
	fmt.Println("")
}
