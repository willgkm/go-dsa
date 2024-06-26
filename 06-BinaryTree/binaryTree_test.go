package main

import (
	"reflect"
	"testing"
)

func setupTree() BinaryTree {

	tree := BinaryTree{}
	tree.insert(4)
	tree.insert(2)
	tree.insert(6)
	tree.insert(1)
	tree.insert(3)
	tree.insert(5)
	tree.insert(7)
	return tree

}

func TestBinaryTree(t *testing.T) {

	t.Run("should create new node in the root ", func(t *testing.T) {

		tree := BinaryTree{}
		tree.insert(1)

		got := tree
		want := BinaryTree{root: &Node{1, nil, nil}}
		assertCorrectMessage(t, got, want)
	})

	t.Run("should insert values in a balance way in a binary tree", func(t *testing.T) {

		tree := setupTree()
		got := tree
		want := BinaryTree{root: &Node{data: 4,
			leftChild:  &Node{data: 2, leftChild: &Node{data: 1}, rightChild: &Node{data: 3}},
			rightChild: &Node{data: 6, leftChild: &Node{data: 5}, rightChild: &Node{data: 7}}}}
		assertCorrectMessage(t, got, want)
	})

	t.Run("should get the height of a 'balance' tree", func(t *testing.T) {

		tree := setupTree()

		got := tree.height()
		want := int8(2)
		assertCorrectMessage(t, got, want)
	})

	t.Run("should get the height of a tree with a bigger right side", func(t *testing.T) {

		tree := setupTree()
		tree.insert(8)
		tree.insert(9)
		tree.insert(10)

		got := tree.height()
		want := int8(5)
		assertCorrectMessage(t, got, want)
	})

	t.Run("should get the inOrderTraversal() of the binary tree", func(t *testing.T) {

		tree := setupTree()

		got := tree.inorderTraversal()
		want := []int8{1, 2, 3, 4, 5, 6, 7}
		assertCorrectMessage(t, got, want)

	})

	t.Run("should get the preOrderTraversal() of the binary tree", func(t *testing.T) {

		tree := setupTree()

		got := tree.preorderTraversal()
		want := []int8{4, 2, 1, 3, 6, 5, 7}
		assertCorrectMessage(t, got, want)
	})

	t.Run("should get the postorderTraversal() of the binary tree", func(t *testing.T) {

		tree := setupTree()

		got := tree.postorderTraversal()
		want := []int8{1, 3, 2, 5, 7, 6, 4}
		assertCorrectMessage(t, got, want)
	})

	t.Run("should get the leverOrderTraversal() of the binary tree", func(t *testing.T) {

		tree := setupTree()

		got := tree.leverOrderTraversal()
		want := []int8{4, 2, 6, 1, 3, 5, 7}
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got any, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
