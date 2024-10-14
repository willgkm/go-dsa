package main

import "fmt"

type Node struct {
	next *Node
	data int8
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) unshift(data int8) {
	// basically create a new node and put in the head of the list
	newNode := &Node{
		data: data,
		next: list.head,
	}
	list.head = newNode
}

func (list *LinkedList) push(data int8) {
	newNode := &Node{data: data}

	// if the head of the list is empty, populate the head with the new node
	if list.head == nil {
		newNode.next = nil
		list.head = newNode
	} else {

		// now if the head is not empty we must run for all values in the list till we find a head with value 'nil'
		current := list.head
		for current.next != nil {
			current = current.next
		}

		current.next = newNode
	}
}

func (list *LinkedList) shift() {
	list.head = list.head.next
}

func (list *LinkedList) pop() {
	if list.head == nil || list.head.next == nil {
		list.head = nil
		return
	}

	current := list.head
	for current.next.next != nil {
		current = current.next
	}

	current.next = nil
}

func (list *LinkedList) display() {
	// for the display, basically we use the seme idea of the insertAtEnd
	// we run all the elements of the list till the current the current is a nil, that means
	// that is the tail or the end of the
	current := list.head
	fmt.Print("head => ")
	for current != nil {
		fmt.Print(current.data, " => ")
		current = current.next
	}
	fmt.Println("tail")
}

func (list *LinkedList) getContent() []int8 {
	arrayList := make([]int8, 0)

	current := list.head
	for current != nil {
		arrayList = append(arrayList, current.data)
		current = current.next
	}
	return arrayList
}

func main() {}
