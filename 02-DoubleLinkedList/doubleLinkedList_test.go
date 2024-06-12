package main

import (
	"reflect"
	"testing"
)

func TestDoubleLinked(t *testing.T) {

	t.Run("should get the content inside a empty double linked list", func(t *testing.T) {
		list := DoubleLinkedList{}

		got := list.getContent()
		want := []int8{}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should get the content inside a not empty double linked list", func(t *testing.T) {
		list := DoubleLinkedList{head: &Node{prev: nil, data: 1, next: nil}}

		got := list.getContent()
		want := []int8{1}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert a data at the end of inside a empty double linked list", func(t *testing.T) {
		list := DoubleLinkedList{}
		list.insertAtEnd(1)
		got := list.getContent()
		want := []int8{1}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert a data at the end of a not empty double linked list", func(t *testing.T) {
		list := DoubleLinkedList{head: &Node{prev: nil, data: 1, next: nil}}
		list.insertAtEnd(2)

		got := list.getContent()
		want := []int8{1, 2}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert a data at the begin of inside a empty double linked list", func(t *testing.T) {
		list := DoubleLinkedList{}
		list.insertAtBegin(1)
		got := list.getContent()
		want := []int8{1}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert a data at the begin of a not empty double linked list", func(t *testing.T) {
		list := DoubleLinkedList{head: &Node{prev: nil, data: 1, next: nil}}
		list.insertAtBegin(2)

		got := list.getContent()
		want := []int8{2, 1}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should remove a specific data in the beginning of a double linked list", func(t *testing.T) {
		list := DoubleLinkedList{}
		list.insertAtEnd(1)
		list.insertAtEnd(2)
		list.insertAtEnd(3)
		list.remove(1)

		got := list.getContent()
		want := []int8{2, 3}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should remove a specific data in the middle of a double linked list", func(t *testing.T) {
		list := DoubleLinkedList{}
		list.insertAtEnd(1)
		list.insertAtEnd(2)
		list.insertAtEnd(3)
		list.remove(2)

		got := list.getContent()
		want := []int8{1, 3}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should remove a specific data in the middle of a double linked list", func(t *testing.T) {
		list := DoubleLinkedList{}
		list.insertAtEnd(1)
		list.insertAtEnd(2)
		list.insertAtEnd(3)
		list.insertAtEnd(4)
		list.insertAtEnd(5)
		list.insertAtEnd(6)
		list.remove(4)

		got := list.getContent()
		want := []int8{1, 2, 3, 5, 6}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should remove a specific data at the end off a double linked list", func(t *testing.T) {
		list := DoubleLinkedList{}
		list.insertAtEnd(1)
		list.insertAtEnd(2)
		list.insertAtEnd(3)
		list.remove(3)

		got := list.getContent()
		want := []int8{1, 2}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should get a lenght of a double linked list", func(t *testing.T) {
		list := DoubleLinkedList{}
		list.insertAtEnd(1)
		list.insertAtEnd(2)
		list.insertAtEnd(3)
		list.insertAtEnd(4)
		list.insertAtEnd(5)
		list.insertAtEnd(6)
		got := list.lenght()
		var want int8 = 5

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got any, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
