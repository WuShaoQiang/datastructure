package datastructure

import "log"

type LinkListNode struct {
	Value interface{}
	Next  *LinkListNode
	Prev  *LinkListNode
}

type LinkList struct {
	len  int
	head *LinkListNode
	tail *LinkListNode
}

func NewLinkList() *LinkList {
	return &LinkList{
		len:  0,
		head: nil,
		tail: nil,
	}
}

func NewLinkListNode(value interface{}) *LinkListNode {
	return &LinkListNode{
		Value: value,
		Next:  nil,
		Prev:  nil,
	}
}

func (l *LinkList) Len() int {
	return l.len
}

func (l *LinkList) IsEmpty() bool {
	return l.len == 0
}

func (l *LinkList) Clear() {
	l.len = 0
	l.head = nil
	l.tail = nil
}

func (l *LinkList) Append(value interface{}) {
	node := NewLinkListNode(value)

	if l.len == 0 {
		l.head = node
		l.tail = node
	} else {
		l.tail.Next = node
		node.Prev = l.tail

		l.tail = node
	}

	l.len++
}

func (l *LinkList) Prepend(value interface{}) {
	node := NewLinkListNode(value)

	if l.len == 0 {
		l.head = node
		l.tail = node
	} else {
		l.head.Prev = node
		node.Next = l.head

		l.head = node
	}

	l.len++
}

func (l *LinkList) Add(value interface{}, index int) {
	if index > l.len+1 {
		log.Println("index too large")
		return
	}

	if index < 1 {
		log.Println("index too small")
		return
	}

	node := NewLinkListNode(value)

	prev := l.Search(index - 1)

	node.Prev = prev
	node.Next = prev.Next

	prev.Next.Prev = node
	prev.Next = node

}

func (l *LinkList) Search(index int) *LinkListNode {
	if index > l.len+1 {
		log.Println("index too large")
		return nil
	}

	if index < 1 {
		log.Println("index too small")
		return nil
	}

	var node *LinkListNode

	if index == 1 {
		return l.head
	}

	if index == l.len {
		return l.tail
	}

	if l.len/2 >= index { //左边至少第2个
		node = l.head
		for i := 1; i < index; i++ {
			node = node.Next
		}
		return node
	} else { //右边至少第2个
		node = l.tail
		for i := l.len - 1; i > index; i-- {
			node = node.Prev
		}
		return node
	}
}
