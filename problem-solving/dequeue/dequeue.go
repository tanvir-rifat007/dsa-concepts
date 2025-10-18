package main

import "fmt"

type Node[T any] struct {
	val  T
	next *Node[T]
	prev *Node[T]
}

type Dequeue[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

func NewDequeue[T any]() *Dequeue[T] {

	return &Dequeue[T]{
		size: 0,
		head: nil,
		tail: nil,
	}

}

func (d *Dequeue[T]) pushFront(val T) {

	node := &Node[T]{
		val: val,
	}

	if d.head == nil {
		d.head = node
		d.tail = node
		d.size++
	} else {
		d.head.prev = node
		node.next = d.head
		d.head = node
		d.size++
	}

}

func (d *Dequeue[T]) pushBack(val T) {
	node := &Node[T]{
		val: val,
	}

	if d.tail == nil {
		d.head = node
		d.tail = node
		d.size++
	} else {
		d.tail.next = node
		node.prev = d.tail
		d.tail = node
		d.size++
	}

}

func (d *Dequeue[T]) print() {
	cur := d.head

	for cur != nil {
		fmt.Printf("%v ", cur.val)
		cur = cur.next

	}
	fmt.Println()

}

func (d *Dequeue[T]) popRight() T {

	if d.tail == nil {
		return *new(T)
	}

	if d.size == 1 {
		d.head = nil
		d.tail = nil
		d.size--

		return *new(T)
	}
	cur := d.tail
	cur.prev.next = nil
	d.tail = cur.prev
	d.size--

	return cur.val

}

func (d *Dequeue[T]) popLeft() T {
	if d.head == nil {
		return *new(T)

	}

	if d.size == 1 {
		d.head = nil
		d.tail = nil
		d.size--

		return *new(T)
	}

	cur := d.head
	cur.next.prev = nil
	d.head = cur.next
	d.size--

	return cur.val

}

func main() {

	dequeue := NewDequeue[string]()

	dequeue.pushFront("200")
	dequeue.pushFront("100")

	dequeue.pushBack("300")

	dequeue.popRight()
	dequeue.popLeft()

	// for size = 1
	dequeue.popLeft()

	dequeue.pushBack("300")
	dequeue.pushFront("200")

	dequeue.pushFront("100")

	dequeue.print()

	fmt.Printf("%d\n", dequeue.size)

}
