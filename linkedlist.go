package main

import "fmt"

type LinkedList[T any] struct {
	head, tail *Node[T]
}

type Node[T any] struct {
	prev *Node[T]
	next *Node[T]
	val  T
}

func (l *LinkedList[T]) Append(el T) {
	if l.tail == nil {
		l.head = &Node[T]{val: el}
		l.tail = l.head
	} else {
		tmp := l.tail
		l.tail.next = &Node[T]{val: el}
		l.tail = l.tail.next
		l.tail.prev = tmp
	}
}

func (l *LinkedList[T]) Prepend(el T) {
	if l.head == nil {
		l.head = &Node[T]{val: el}
		l.tail = l.head
	} else {
		tmp := l.head
		l.head = &Node[T]{val: el}
		l.head.next = tmp
		l.head.next.prev = l.head
	}
}

func (lst *LinkedList[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	list := LinkedList[int]{}
	list.Append(10)
	list.Append(2)
	list.Append(550)
	list.Append(2)
	list.Prepend(69)

	fmt.Println(list.head.next.next.next.prev)

	fmt.Println(list.GetAll())
}
