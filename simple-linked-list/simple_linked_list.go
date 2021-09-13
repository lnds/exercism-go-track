package linkedlist

import "fmt"

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func NewElement(data int) *Element {
	return &Element{
		data: data,
		next: nil,
	}
}

func New(elements []int) *List {
	list := &List{head: nil, size: 0}
	for _, e := range elements {
		list.Push(e)
	}
	return list
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Push(data int) {
	if list.head == nil {
		list.head = NewElement(data)
		list.size = 1
	} else {
		p := list.head
		for p.next != nil {
			p = p.next
		}
		p.next = NewElement(data)
		list.size++
	}
}

func (list *List) Pop() (int, error) {
	if list.head == nil {
		return 0, fmt.Errorf("empty list")
	}
	p := list.head
	prv := list.head
	for p.next != nil {
		prv = p
		p = p.next
	}
	prv.next = nil
	list.size--
	return p.data, nil
}

func (list *List) Array() []int {
	if list == nil {
		return nil
	}
	result := []int{}
	p := list.head
	for p != nil {
		result = append(result, p.data)
		p = p.next
	}
	return result
}

func (list *List) Reverse() *List {
	if list == nil {
		return nil
	}
	newList := New(nil)
	for list.Size() > 0 {
		if data, err := list.Pop(); err == nil {
			newList.Push(data)
		}
	}
	return newList
}
