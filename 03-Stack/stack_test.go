package main

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {

	t.Run("should insert an element into the stack", func(t *testing.T) {
		stack := Stack{}
		stack.push(1)

		got := stack.values
		want := []int8{1}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert more than an element into the stack", func(t *testing.T) {
		stack := Stack{}
		stack.push(1)
		stack.push(2)
		stack.push(3)
		stack.push(4)

		got := stack.values
		want := []int8{1, 2, 3, 4}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should remove an element into the stack", func(t *testing.T) {
		stack := Stack{values: []int8{1}}
		stack.pop()

		got := stack.values
		want := []int8{}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should remove more than an element into the stack", func(t *testing.T) {
		stack := Stack{values: []int8{1, 2, 3, 4}}
		stack.pop()
		stack.pop()

		got := stack.values
		want := []int8{1, 2}

		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert the element at top of a stack", func(t *testing.T) {
		stack := Stack{values: []int8{1, 2, 3, 4}}

		got := stack.top()
		want := int8(4)

		assertCorrectMessage(t, got, want)
	})

	t.Run("should verify the empty stack", func(t *testing.T) {
		stack := Stack{values: []int8{}}

		got := stack.isEmpty()
		want := true

		assertCorrectMessage(t, got, want)
	})

	t.Run("should verify a not empty stack", func(t *testing.T) {
		stack := Stack{values: []int8{1, 2, 3, 4}}

		got := stack.isEmpty()
		want := false

		assertCorrectMessage(t, got, want)
	})

	t.Run("should return a size of a empty stack", func(t *testing.T) {
		stack := Stack{values: []int8{}}

		got := stack.size()
		want := int8(0)

		assertCorrectMessage(t, got, want)
	})

	t.Run("should return a size of an stack with values", func(t *testing.T) {
		stack := Stack{values: []int8{1, 2, 3, 4}}

		got := stack.size()
		want := int8(4)

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got any, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
