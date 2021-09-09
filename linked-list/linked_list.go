package linkedlist

import (
	"errors"
)

type Node struct {
	prev *Node
	next *Node
	Val  interface{}
}

func NewNode(val interface{}) *Node {
	return &Node{
		Val: val,
	}
}

func (node *Node) Next() *Node {
	return node.next
}

func (node *Node) Prev() *Node {
	return node.prev
}

type List struct {
	head *Node
	tail *Node
}

var ErrEmptyList = errors.New("empty list")

func NewList(args ...interface{}) *List {
	list := &List{
		head: nil,
		tail: nil,
	}
	for _, arg := range args {
		list.PushBack(arg)
	}
	return list
}

func (list *List) First() *Node {
	return list.head
}

func (list *List) Last() *Node {
	return list.tail
}

func (list *List) PushFront(v interface{}) {
	node := NewNode(v)
	if list.head == nil {
		list.tail = node
	} else {
		node.next = list.head
		list.head.prev = node
	}
	list.head = node
}

func (list *List) PopFront() (interface{}, error) {
	if list.head == nil {
		return nil, ErrEmptyList
	}
	v := list.head.Val
	list.head = list.head.next
	if list.head == nil {
		list.tail = nil
	} else {
		list.head.prev = nil
	}
	return v, nil
}

func (list *List) PushBack(v interface{}) {
	node := NewNode(v)
	if list.tail == nil {
		list.head = node
	} else {
		node.prev = list.tail
		list.tail.next = node
	}
	list.tail = node

}

func (list *List) PopBack() (interface{}, error) {
	if list.tail == nil {
		return 0, ErrEmptyList
	}
	v := list.tail.Val
	list.tail = list.tail.prev
	if list.tail == nil {
		list.head = nil
	} else {
		list.tail.next = nil
	}
	return v, nil
}

func (list *List) Reverse() {
	for e := list.First(); e != nil; e = e.Prev() {
		e.next, e.prev = e.prev, e.next
	}
	list.head, list.tail = list.tail, list.head
}
