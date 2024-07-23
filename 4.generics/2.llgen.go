package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

type Linkedlist[T any] struct {
	head *List[T]
}

func newlist[T any]() *Linkedlist[T] {
	return &Linkedlist[T]{}
}

func (ll *Linkedlist[T]) append(data T) {

	newnode := List[T]{next: nil, val: data}

	if ll.head == nil {
		ll.head = &newnode
		return
	}
	curr := ll.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = &newnode
}

func (ll *Linkedlist[T]) show() {
	curr := ll.head

	for curr != nil {
		fmt.Print(curr.val , " -> ")
		curr = curr.next
	}
	fmt.Println("nil")
}

type kumkum struct {
	name string
	addr string
}

func main() {

	kk := kumkum{"kum","kum"}

	nll := newlist[any]()

	nll.append(10)
	nll.append(string('$'))
	nll.append(true)
	nll.append("vai da dei")
	nll.append(20.34)
	nll.append(kk)

	nll.show()

}
