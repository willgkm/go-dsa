package main

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {

	t.Run("get a empty array from getContent", func(t *testing.T) {

		list := LinkedList{}
		got := list.getContent()
		want := make([]int8, 0)
		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert a value in a empty linkedList using push func", func(t *testing.T) {

		list := LinkedList{}
		list.push(1)
		got := list.getContent()
		want := []int8{1}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert a value at the end of linkedList using push func", func(t *testing.T) {

		list := LinkedList{&Node{data: 1, next: nil}}
		list.push(2)
		got := list.getContent()
		want := []int8{1, 2}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert a value on a empty LinkedList using unshift func", func(t *testing.T) {

		list := LinkedList{}
		list.unshift(0)
		got := list.getContent()
		want := []int8{0}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert the value as first item of the LinkedList", func(t *testing.T) {

		list := LinkedList{&Node{data: 1, next: &Node{data: 2, next: nil}}}
		list.unshift(0)
		got := list.getContent()
		want := []int8{0, 1, 2}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should delete a value at the begin of a linkedList who have 3 values", func(t *testing.T) {

		list := LinkedList{&Node{data: 1, next: &Node{data: 2, next: &Node{data: 3, next: nil}}}}
		list.shift()
		got := list.getContent()
		want := []int8{2, 3}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should delete a value at the end of a linkedList who have 3 values", func(t *testing.T) {

		list := LinkedList{&Node{data: 1, next: &Node{data: 2, next: &Node{data: 3, next: nil}}}}
		list.pop()
		got := list.getContent()
		want := []int8{1, 2}

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got any, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
