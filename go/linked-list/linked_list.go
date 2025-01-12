package linkedlist

import "fmt"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.

type Node struct {
	Value int
	next  *Node
	prev  *Node
}

type List struct {
	root *Node
}

func NewList(elements ...interface{}) *List {
	list := &List{
		root: nil,
	}

	if len(elements) == 0 {
		return list
	}

	var prev *Node = nil
	var current *Node = nil
	var root *Node = nil
	for _, el := range elements {
		val := el.(int)
		current = &Node{
			Value: val,
			prev:  prev,
			next:  nil,
		}
		if prev != nil {
			prev.next = current
		}
		prev = current
		if root == nil {
			root = current
		}
	}

	list.root = root
	return list

}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	prevRoot := l.root
	val := v.(int)
	node := &Node{
		Value: val,
		next:  nil,
		prev:  nil,
	}
	l.root = node
	l.root.next = prevRoot
	if prevRoot != nil {
		prevRoot.prev = l.root
	}
}

func (l *List) Push(v interface{}) {
	val := v.(int)
	if l.root == nil {
		l.root = &Node{
			Value: val,
			next:  nil,
		}
		return
	}
	current := l.root
	for current.next != nil {
		current = current.next
	}
	next := &Node{
		Value: val,
		next:  nil,
		prev:  current,
	}
	current.next = next
}

func (l *List) Shift() (interface{}, error) {
	if l.root == nil {
		return nil, fmt.Errorf("shift operation on an empty linked list")
	}
	val := l.root.Value
	if l.root.next != nil {
		l.root.next.prev = nil
	}
	next := l.root.next
	l.root.next = nil
	l.root.prev = nil
	l.root = next
	return val, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.root == nil {
		return nil, fmt.Errorf("cannot pop an empty linked list")
	}
	if l.root.next == nil {
		val := l.root.Value
		l.root = nil
		return val, nil
	}
	current := l.root
	var prev *Node = nil

	for current.next != nil {
		prev = current
		current = current.next
	}
	result := current
	prev.next = nil
	current.prev = nil

	return result.Value, nil

}

func (l *List) Reverse() {
	if l.root == nil {
		return
	}
	current := l.root
	var prev *Node = nil

	for current != nil {
		next := current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}
	l.root = prev
}

func (l *List) First() *Node {
	return l.root
}

func (l *List) Last() *Node {
	if l.root == nil {
		return nil
	}
	current := l.root
	for current.next != nil {
		current = current.next
	}
	return current
}
