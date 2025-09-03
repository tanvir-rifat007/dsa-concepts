package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type SinglyLinkedList struct {
	size int
	head *Node
}

func (this *SinglyLinkedList) push(val int) {

	node := &Node{
		val:  val,
		next: nil,
	}

	if this.size == 0 {
		this.head = node
		this.size++
		return

	}

	curr := this.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = node
	this.size++

}

func (this *SinglyLinkedList) pop() int {
	if this.size == 0 {

		return -1
	}
	if this.size == 1 {
		val := this.head.val
		this.head = nil
		this.size--
		return val
	}

	curr := this.head

	for curr.next.next != nil {
		curr = curr.next
	}
	val := curr.next.val
	curr.next = nil
	this.size--
	return val

}

func (this *SinglyLinkedList) peek() int {
	// return the value of the head node without removing it

	if this.size == 0 {
		return -1

	}
	if this.size == 1 {
		return this.head.val
	}

	curr := this.head

	for curr.next != nil {
		curr = curr.next
	}
	return curr.val

}

func (this *SinglyLinkedList) print() {
	curr := this.head

	for curr != nil {
		fmt.Println(curr.val)
		curr = curr.next

	}

}

func main() {

	sll := &SinglyLinkedList{}
	sll.push(1)
	sll.push(2)
	sll.push(3)
	sll.push(4)

	fmt.Println("Popped:", sll.pop())
	fmt.Println("Popped:", sll.pop())

	fmt.Println("Peek:", sll.peek())

	sll.print()
}
