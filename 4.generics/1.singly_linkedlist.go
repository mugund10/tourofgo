package main

import "fmt"

type nodee struct {
	data int
	next *nodee
}

type linkedlist struct {
	head *nodee
}

func newll() *linkedlist {
	return &linkedlist{}
}

// add node at end
func (ll *linkedlist) append(data int) {
	newnode := &nodee{data: data, next: nil}
	//check whether the head is empty or not
	if ll.head == nil {
		ll.head = newnode
		return
	}
	//iterate over all nodes until the last nodes gets selected
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	//adding the node at last
	current.next = newnode
}

func (ll *linkedlist) display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := newll()
	list.append(10)
	list.append(34)
	list.append(56)

	list.display()
}
