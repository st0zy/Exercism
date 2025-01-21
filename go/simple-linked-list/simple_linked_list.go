package linkedlist

import "errors"

// Define the List and Element types here.

type List struct {
	head *Node
	size int
}

type Node struct {
	data int
	next *Node
}

func New(elements []int) *List {
	var result = &List{}
	var cur *Node
	var prev *Node
	for _, el := range elements {
		cur = &Node{data: el}
		if prev == nil {
			result.head = cur
		} else {
			prev.next = cur
		}
		result.size++
		prev = cur
	}
	return result
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	if l.head == nil {
		l.head = &Node{
			data: element,
			next: nil,
		}
		l.size = 1
		return
	}
	current := l.head
	prev := l.head
	for ; current != nil; current = current.next {
		prev = current
	}
	prev.next = &Node{
		data: element,
		next: nil,
	}
	l.size++

}

func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("can't pop from an empty list")
	}

	var current = l.head
	var prev *Node
	for ; current != nil && current.next != nil; current = current.next {
		prev = current
	}
	if prev != nil {
		prev.next = nil
	}
	l.size--

	return current.data, nil
}

func (l *List) Array() []int {
	result := make([]int, 0)
	if l == nil {
		return nil
	}
	var curr *Node = l.head

	for ; curr != nil; curr = curr.next {
		result = append(result, curr.data)
	}
	return result
}

func (l *List) Reverse() *List {

	var curr *Node = l.head
	var prev *Node
	var next *Node

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	return &List{
		head: prev,
		size: l.size,
	}
}
