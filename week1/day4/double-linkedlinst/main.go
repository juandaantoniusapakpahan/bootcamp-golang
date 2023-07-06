package main

import "fmt"

type DoubleLinked interface {
	Add(val int)
	Remove(val int)
	ShowAll()
	DeleteByVal(val int)
}

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type List struct {
	head *Node
}

func NewList() DoubleLinked {
	return &List{}
}

func (l *List) DeleteByVal(val int) {

	if l.head == nil {
		fmt.Println("There is no data")
		return
	}

	if l.head.Val == val {
		l.head = l.head.Next
		l.head.Prev = nil
	}

	cur := l.head
	for cur.Next != nil {
		if cur.Next.Val != val {
			cur = cur.Next
		} else {
			cur.Next = cur.Next.Next
			cur.Next.Prev = cur
		}
	}
	if cur.Val == val {
		cur.Prev.Next = nil
		cur.Prev = nil
	}

}

func (l *List) ShowAll() {
	if l.head == nil {
		fmt.Println("There is not data")
		return
	}

	cur := l.head
	for cur != nil {
		fmt.Println(*&cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

func (l *List) Add(val int) {
	newNode := Node{Val: val}

	if l.head == nil {
		l.head = &newNode
		return
	}

	cur := l.head
	for cur.Next != nil {
		cur = cur.Next
	}

	if cur.Next == nil {
		cur.Next = &newNode
		cur.Next.Prev = cur
	}
}
func (l *List) Remove(val int) {
	if l.head == nil {
		fmt.Println("There is no data")
		return
	}

	if l.head.Val == val {
		l.head = l.head.Next
		l.head.Prev = nil
		return
	}

	cur := l.head
	for cur.Next != nil && cur.Next.Val != val {
		cur = cur.Next
	}

	cur.Next = cur.Next.Next
	cur.Next.Prev = cur
}

func main() {
	list := NewList()

	list.Add(3)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(3)
	list.Add(4)

	list.ShowAll()
	// list.Remove(2)
	// list.ShowAll()
	list.DeleteByVal(3)
	list.ShowAll()

}
