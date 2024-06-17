package main

import (
	"reflect"
	"testing"
)

func TestCircularLinkedList(t *testing.T) {
	t.Run("should create a node with just a value", func(t *testing.T) {

		got := createNode(1, nil)
		want := &Node{value: 1, next: nil}
		assertCorrectMessage(t, got, want)
	})

	t.Run("should create a node with value and a next", func(t *testing.T) {
		node := createNode(2, nil)
		got := createNode(1, node)
		want := &Node{value: 1, next: &Node{value: 2, next: nil}}
		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert a data in the begin of a empty list", func(t *testing.T) {
		list := CircularLinkedList{}
		list.insertAtBegin(1)

		var node *Node = &Node{value: 1, next: nil}
		node.next = node
		var expectedList = CircularLinkedList{head: node, last: node}

		got := list
		want := expectedList
		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert a data in the begin of a list", func(t *testing.T) {
		list := CircularLinkedList{}
		list.insertAtBegin(2)
		list.insertAtBegin(1)
		list.insertAtBegin(0)

		var node1 *Node = &Node{value: 2, next: nil}
		var node2 *Node = &Node{value: 1, next: nil}
		var node3 *Node = &Node{value: 0, next: nil}
		node3.next = node2
		node2.next = node1
		node1.next = node3

		var expectedList = CircularLinkedList{head: node3, last: node1}
		expectedList.head = node3
		expectedList.last = node1

		got := list
		want := expectedList
		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert a data in the end of a empty list", func(t *testing.T) {
		list := CircularLinkedList{}
		list.insertAtEnd(1)

		var node *Node = &Node{value: 1, next: nil}
		node.next = node
		var expectedList = CircularLinkedList{head: node, last: node}

		got := list
		want := expectedList
		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert a data in the end of a list", func(t *testing.T) {
		list := CircularLinkedList{}
		list.insertAtEnd(1)
		list.insertAtEnd(2)
		list.insertAtEnd(3)

		var node1 *Node = &Node{value: 1, next: nil}
		var node2 *Node = &Node{value: 2, next: nil}
		var node3 *Node = &Node{value: 3, next: nil}
		node1.next = node2
		node2.next = node3
		node3.next = node1

		var expectedList = CircularLinkedList{head: node3, last: node3}
		expectedList.head = node1
		expectedList.last = node3

		got := list
		want := expectedList
		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert a data in between 2 nodes", func(t *testing.T) {
		list := CircularLinkedList{}
		list.insertAtEnd(1)
		list.insertAtEnd(3)
		list.insertAfter(2, 1)

		var node1 *Node = &Node{value: 1, next: nil}
		var node2 *Node = &Node{value: 2, next: nil}
		var node3 *Node = &Node{value: 3, next: nil}
		node1.next = node2
		node2.next = node3
		node3.next = node1

		var expectedList = CircularLinkedList{head: node3, last: node3}
		expectedList.head = node1
		expectedList.last = node3

		got := list
		want := expectedList
		assertCorrectMessage(t, got, want)
	})
	t.Run("should delete a node in the start of a list ", func(t *testing.T) {
		list := CircularLinkedList{}
		list.insertAtEnd(0)
		list.insertAtEnd(1)
		list.insertAtEnd(2)
		list.insertAtEnd(3)
		list.removeFirst()

		var node1 *Node = &Node{value: 1, next: nil}
		var node2 *Node = &Node{value: 2, next: nil}
		var node3 *Node = &Node{value: 3, next: nil}
		node1.next = node2
		node2.next = node3
		node3.next = node1

		var expectedList = CircularLinkedList{head: node3, last: node3}
		expectedList.head = node1
		expectedList.last = node3

		got := list
		want := expectedList
		assertCorrectMessage(t, got, want)
	})

	t.Run("should delete a node in the end of a list ", func(t *testing.T) {
		list := CircularLinkedList{}
		list.insertAtEnd(1)
		list.insertAtEnd(2)
		list.insertAtEnd(3)
		list.insertAtEnd(4)
		list.removeLast()

		var node1 *Node = &Node{value: 1, next: nil}
		var node2 *Node = &Node{value: 2, next: nil}
		var node3 *Node = &Node{value: 3, next: nil}
		node1.next = node2
		node2.next = node3
		node3.next = node1

		var expectedList = CircularLinkedList{head: node3, last: node3}
		expectedList.head = node1
		expectedList.last = node3

		got := list
		want := expectedList
		assertCorrectMessage(t, got, want)
	})

	t.Run("should delete a node in the any value of the list ", func(t *testing.T) {
		list := CircularLinkedList{}
		list.insertAtEnd(1)
		list.insertAtEnd(2)
		list.insertAtEnd(5)
		list.insertAtEnd(3)
		list.remove(5)

		var node1 *Node = &Node{value: 1, next: nil}
		var node2 *Node = &Node{value: 2, next: nil}
		var node3 *Node = &Node{value: 3, next: nil}
		node1.next = node2
		node2.next = node3
		node3.next = node1

		var expectedList = CircularLinkedList{head: node3, last: node3}
		expectedList.head = node1
		expectedList.last = node3

		got := list
		want := expectedList
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got any, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
