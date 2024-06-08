package main

import "fmt"

type Node struct {
	prev *Node
	data int8
	next *Node
}

type DoubleLinkedList struct {
	head *Node
}

func (list *DoubleLinkedList) unshift(data int8) {
	newNode := &Node{data: data}

	newNode.next = list.head
	newNode.prev = nil

	if list.head != nil {
		list.head.prev = newNode
	}

	list.head = newNode

}

func (list *DoubleLinkedList) push() {

}

func (list *DoubleLinkedList) displayForward() {

	current := list.head

	fmt.Print("head => ")
	for current != nil {
		fmt.Print(current.data, " => ")
		current = current.next
	}
	fmt.Println("tail")

}

func (list *DoubleLinkedList) getContent() []int8 {

	current := list.head
	arrayList := make([]int8, 0)
	for current != nil {
		arrayList = append(arrayList, current.data)
		current = current.next
	}

	return arrayList
}

func main() {}
