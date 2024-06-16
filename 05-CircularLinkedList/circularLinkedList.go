package main

type Node struct {
	value int8
	next  *Node
}

type CircularLinkedList struct {
	head *Node
	last *Node
}

func createNode(value int8, next *Node) *Node {
	return &Node{value: value, next: next}
}

func insertOnEmptyList(list *CircularLinkedList, data int8) {
	list.head = createNode(data, nil)
	list.head.next = list.head
	list.last = list.head
}

func (list *CircularLinkedList) insertAtBegin(data int8) {

	if list.head == nil {
		//insert in the begin empty list
		insertOnEmptyList(list, data)

	} else {
		//insert in the begin with values
		newNode := createNode(data, list.head)

		list.head = newNode
		list.last.next = newNode
	}

}

func (list *CircularLinkedList) insertAtEnd(data int8) {
	if list.head == nil {
		insertOnEmptyList(list, data)
	} else {
		newNode := createNode(data, list.head)
		list.last.next = newNode

		list.last = newNode
	}
}

func (list *CircularLinkedList) insertAfter(data int8, node int8) {
	if list.head == nil {
		insertOnEmptyList(list, data)
	} else {
		current := list.head
		for current.value != node {
			current = current.next
		}
		newNode := createNode(data, current.next)
		current.next = newNode
	}

}

func (list *CircularLinkedList) remove(data int8) {

}

func main() {}
