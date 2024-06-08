package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {

	t.Run("should get the content inside a double linked list", func(t *testing.T) {
		list := DoubleLinkedList{head: &Node{prev: nil, data: 1, next: nil}}

		got := list.getContent()
		want := []int8{1}

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got any, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
