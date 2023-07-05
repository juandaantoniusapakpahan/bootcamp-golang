package main

import (
	"fmt"
)

type Interface interface {
	Add(val int)
	Remove(val int)
	TolongPrint()
	RemoveBySame(val int)
}

type LinkList struct {
	Val  int
	Next *LinkList
}

type List struct {
	head *LinkList
}

func (l *List) Add(val int) {
	node := &LinkList{Val: val}
	if l.head == nil {
		l.head = node
		return
	}

	cur := l.head
	for cur.Next != nil {
		cur = cur.Next
	}

	if cur.Next == nil {
		cur.Next = node
	}

}

func (l *List) Remove(val int) {
	if l.head == nil {
		return
	}

	if l.head.Val == val {
		l.head = l.head.Next
		return
	}

	cur := l.head
	for cur.Next != nil && cur.Next.Val != val {
		cur = cur.Next
	}

	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
}

func (l *List) RemoveBySame(val int) {

	var (
		rear = l.head
		pre  = new(LinkList)
		cur  = pre
	)

	pre.Next = l.head

	for rear != nil {
		if rear.Val == val {
			cur.Next = cur.Next.Next
			rear = cur.Next

		} else {
			cur = rear
			rear = rear.Next
		}

	}
	pre.Next = pre.Next.Next

}
func (l *List) TolongPrint() {
	if l.head == nil {
		fmt.Println(l.head)
	}
	cur := l.head
	for cur != nil {
		fmt.Printf("%d \n", cur.Val)
		cur = cur.Next
	}
}

func main() {

	var list Interface = &List{}

	list.Add(4)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(1)
	list.Add(4)

	list.TolongPrint()
	fmt.Println()

	list.RemoveBySame(4)
	fmt.Println()
	list.TolongPrint()

}
