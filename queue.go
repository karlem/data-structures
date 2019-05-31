package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	front *Node
	back  *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(data interface{}) {
	newNode := &Node{data, nil}

	if q.front == nil {
		q.front = newNode
		q.back = newNode
	} else {
		q.back.next = newNode
		q.back = newNode
	}
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.front == nil {
		return nil, errors.New("Queue is empty")
	}

	out := q.front
	q.front = out.next

	if q.front == nil {
		q.back = nil
	}

	return out.data, nil
}

func (q *Queue) Print() {
	if q.front == nil {
		fmt.Println("Queue is empty")
	}

	node := q.front
	for node != nil {
		fmt.Printf("=> %v ", node.data)
		node = node.next
	}
	fmt.Println("")
}
