package main

func main() {
	lil := NewLinkedList()
	lil.addAtFront(1)
	lil.addAtFront(2)
	lil.addAtFront(4)
	lil.addAtEnd(5)

	lil.print()
}
