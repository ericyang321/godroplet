package linkedlist

import (
	"fmt"
	"strings"
)

// Node is single data container for LinkedList
type Node struct {
	value interface{}
	next  *Node
}

// List is a linked list Node holder
type List struct {
	head *Node
}

// ErrLinkedList is error type wrapper for this linked list.
type ErrLinkedList struct {
	message string
}

func (e *ErrLinkedList) Error() string {
	return e.message
}

// NewErrLinkedList instantiates new error for an empty list.
func NewErrLinkedList(message string) *ErrLinkedList {
	return &ErrLinkedList{
		message: message,
	}
}

// Get fetches first node containing param value.
func (L *List) Get(val interface{}) (*Node, *ErrLinkedList) {
	if L.head == nil {
		return nil, NewErrLinkedList("List is empty.")
	}
	n := L.head
	for n != nil {
		if n.value == val {
			return n, nil
		}
		n = n.next
	}
	return nil, NewErrLinkedList("Linked list doesn't have the value you're searching for.")
}

// GetIndex fetches the node on input index
func (L *List) GetIndex(i int) (*Node, *ErrLinkedList) {
	if L.head == nil {
		return nil, NewErrLinkedList("List is empty.")
	}
	n := L.head
	for i > 0 {
		if n != nil {
			n = n.next
			i--
		} else {
			return nil, NewErrLinkedList("Index overflow.")
		}
	}
	return n, nil
}

// Insert creates a Node and appends it to linked list
func (L *List) Insert(val interface{}) *Node {
	newNode := &Node{
		value: val,
		next:  nil,
	}
	if L.head == nil {
		L.head = newNode
		return newNode
	}
	n := L.head
	for n.next != nil {
		n = n.next
	}
	n.next = newNode
	return newNode
}

// Pop removes the tail most node of the linkedlist.
func (L *List) Pop() (*Node, *ErrLinkedList) {
	if L.head == nil {
		return nil, NewErrLinkedList("List is empty.")
	}
	f := L.head
	g := f.next
	if g == nil {
		L.head = nil
		return f, nil
	}
	for g != nil && g.next != nil {
		f = f.next
		g = g.next
	}
	f.next = nil
	return g, nil
}

// Print prints out all values inside the linked list, in order from head to tail
func (L *List) Print() string {
	if L.head == nil {
		return "[]"
	}
	buffer := make([]string, 0)
	buffer = append(buffer, fmt.Sprintf("[HEAD: %v] -> ", L.head.value))
	n := L.head.next
	for n != nil {
		buffer = append(buffer, fmt.Sprintf("[%v] -> ", n.value))
		n = n.next
	}
	str := strings.Join(buffer, "")
	fmt.Println(str)
	return str
}
