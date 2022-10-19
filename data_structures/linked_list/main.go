package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type List struct {
	Head *Node
	len  int
}

//create a new list
func newList() *List {
	return &List{
		Head: nil,
		len:  0,
	}
}

//insert value at end of list
func (l *List) Insert(val int) {

	n := Node{}
	n.Val = val

	if l.Head == nil {
		l.Head = &n
		l.len++
		return
	}

	pt := l.Head
	for i := 1; i <= l.len; i++ {
		if pt.Next == nil {
			pt.Next = &n
			l.len++
			return
		}
		pt = pt.Next
	}
}

//display a list
func (l *List) Display() {
	if l.Head == nil {
		fmt.Println("No nodes in list!")
	}

	pt := l.Head
	for i := 1; i <= l.len; i++ {
		fmt.Println("Node Value: ", pt.Val)
		pt = pt.Next
	}
}

func main() {
	ls := newList()

	for i := 1; i <= 10; i++ {
		v := i % 2

		ls.Insert(v)
	}

	ls.Display()
}
