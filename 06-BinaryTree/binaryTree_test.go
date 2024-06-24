package main

import (
	"reflect"
	"testing"
)

func TestBinaryTree(t *testing.T) {

	t.Run("should create new node in the root ", func(t *testing.T) {

		tree := BinaryTree{}
		tree.insert(1)

		got := tree
		want := BinaryTree{root: &Node{1, nil, nil}}
		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert values in a balance way in a binary tree", func(t *testing.T) {

		tree := BinaryTree{}
		tree.insert(4)
		tree.insert(2)
		tree.insert(6)
		tree.insert(1)
		tree.insert(3)
		tree.insert(5)
		tree.insert(7)
		got := tree

		want := BinaryTree{root: &Node{data: 4,
			leftChild:  &Node{data: 2, leftChild: &Node{data: 1}, rightChild: &Node{data: 3}},
			rightChild: &Node{data: 6, leftChild: &Node{data: 5}, rightChild: &Node{data: 7}}}}
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got any, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}