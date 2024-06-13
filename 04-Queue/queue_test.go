package main

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {

	t.Run("should create a Queue", func(t *testing.T) {
		got := newQueue(0)
		want := &Queue{make([]int8, 0)}
		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert a value on a queue", func(t *testing.T) {
		queue := &Queue{make([]int8, 0)}
		queue.enqueue(1)

		got := []int8{1}
		want := queue.values
		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert a lot of a values on the queue", func(t *testing.T) {
		queue := &Queue{make([]int8, 0)}
		queue.enqueue(1)
		queue.enqueue(2)
		queue.enqueue(3)

		got := queue.values
		want := []int8{1, 2, 3}
		assertCorrectMessage(t, got, want)
	})

	t.Run("should unqueue a value from the queue", func(t *testing.T) {
		queue := &Queue{values: []int8{1, 2, 3}}
		queue.dequeue()
		got := queue.values
		want := []int8{2, 3}
		assertCorrectMessage(t, got, want)
	})

	t.Run("should not return an error from trying enqueue a empty queue", func(t *testing.T) {
		queue := &Queue{values: []int8{}}
		queue.dequeue()
		got := queue.values
		want := []int8{}
		assertCorrectMessage(t, got, want)
	})

	t.Run("should get the front element from a queue", func(t *testing.T) {
		queue := &Queue{values: []int8{1, 2, 3}}

		gotValue, gotValidation := queue.front()
		wantValue, WantValidation := int8(1), true
		assertCorrectMessage(t, gotValue, wantValue)
		assertCorrectMessage(t, gotValidation, WantValidation)
	})

	t.Run("should return an valid return(0) from trying get the front a empty queue", func(t *testing.T) {
		queue := &Queue{values: []int8{0}}
		gotValue, gotValidation := queue.front()
		wantValue, WantValidation := int8(0), true
		assertCorrectMessage(t, gotValue, wantValue)
		assertCorrectMessage(t, gotValidation, WantValidation)
	})

	t.Run("should return an invalid return from trying get the front a empty queue", func(t *testing.T) {
		queue := &Queue{values: []int8{}}
		gotValue, gotValidation := queue.front()
		wantValue, WantValidation := int8(0), false
		assertCorrectMessage(t, gotValue, wantValue)
		assertCorrectMessage(t, gotValidation, WantValidation)
	})

	t.Run("should get the back element from a queue", func(t *testing.T) {
		queue := &Queue{values: []int8{1, 2, 3}}

		gotValue, gotValidation := queue.back()
		wantValue, WantValidation := int8(3), true
		assertCorrectMessage(t, gotValue, wantValue)
		assertCorrectMessage(t, gotValidation, WantValidation)
	})

	t.Run("should return an valid return(0) from trying get the back a empty queue", func(t *testing.T) {
		queue := &Queue{values: []int8{3,2,1,0}}
		gotValue, gotValidation := queue.back()
		wantValue, WantValidation := int8(0), true
		assertCorrectMessage(t, gotValue, wantValue)
		assertCorrectMessage(t, gotValidation, WantValidation)
	})

	t.Run("should return an invalid return from trying get the back a empty queue", func(t *testing.T) {
		queue := &Queue{values: []int8{}}
		gotValue, gotValidation := queue.back()
		wantValue, WantValidation := int8(0), false
		assertCorrectMessage(t, gotValue, wantValue)
		assertCorrectMessage(t, gotValidation, WantValidation)
	})

}

func assertCorrectMessage(t testing.TB, got any, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
