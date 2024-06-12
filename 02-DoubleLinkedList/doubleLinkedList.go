package main

type Node struct {
	prev *Node
	data int8
	next *Node
}

type DoubleLinkedList struct {
	head *Node
}

func (list *DoubleLinkedList) insertAtBegin(data int8) {
	newNode := &Node{data: data}

	newNode.next = list.head
	newNode.prev = nil

	if list.head != nil {
		list.head.prev = newNode
	}

	list.head = newNode

}

func (list *DoubleLinkedList) insertAtEnd(data int8) {

	current := list.head

	if list.head == nil {
		list.head = &Node{data: data}
	} else {
		for current.next != nil {
			current = current.next
		}
		current.next = &Node{prev: current, data: data, next: nil}
	}
}

func (list *DoubleLinkedList) remove(data int8) {
	current := list.head

	for current.data != data {
		current = current.next
	}

	if current.prev == nil {
		//remove the first item of the list
		list.head = current.next
	} else if current.next == nil {
		//remove the last item of the list
		current.prev.next = nil
		current.prev = nil
	} else {
		current.prev.next = current.next
		current.next.prev = current.prev
	}

}

func (list *DoubleLinkedList) lenght() int8 {
	current := list.head
	var lengthCount int8 = 0
	for current.next != nil {
		lengthCount++
		current = current.next
	}

	return lengthCount

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
