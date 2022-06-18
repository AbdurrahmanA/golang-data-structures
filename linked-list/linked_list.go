package main

import "fmt"

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

type List struct {
	Len  uint
	Head *Node
	Mid  *Node
	Tail *Node
}

func NewLinkedList() *List {
	return &List{}
}

func (l *List) AddToHead(value int) {
	newNode := &Node{Value: value}

	if l.Len == 0 {
		l.Head = newNode
		l.Mid = newNode
		l.Tail = newNode
	} else {
		temp := l.Head
		l.Head = newNode
		l.Head.Next = temp
		temp.Prev = l.Head

		if (l.Len+1)%2 == 0 {
			l.Mid = l.Mid.Next
		}
	}
	l.Len++
}

func (l *List) RemoveFromHead() *Node {
	if l.Head == nil {
		return nil
	}

	head := l.Head
	l.Head = l.Head.Next
	l.Head.Prev = nil

	return head
}

func (l *List) AddToTail(value int) {
	newNode := &Node{Value: value}

	if l.Len == 0 {
		l.Head = newNode
		l.Mid = newNode
		l.Tail = newNode
	} else {
		oldTail := l.Tail
		oldTail.Next = newNode
		l.Tail = newNode
		l.Tail.Prev = oldTail

		if (l.Len+1)%2 == 0 {
			l.Mid = l.Mid.Next
		}
	}
	l.Len++
}

func (l *List) RemoveFromTail() *Node {
	if l.Len == 0 {
		return nil
	}

	tail := l.Tail
	l.Tail = tail.Prev
	l.Tail.Next = nil

	return tail
}

func (l *List) Show() {
	if l.Head == nil {
		return
	}

	node := l.Head
	for node != nil {
		fmt.Printf("%p -> %+v \n", node, node)
		node = node.Next
	}
	fmt.Println()
}
