package main

type Node struct {
	data                  int8
	leftChild, rightChild *Node
}

type BinaryTree struct {
	root *Node
}

func createNode(data int8) *Node {
	return &Node{data: data, rightChild: nil, leftChild: nil}
}

func insertAtRoot(tree *BinaryTree, data int8) {
	newNode := createNode(data)
	tree.root = newNode
}

func (tree *BinaryTree) insert(data int8) {
	if tree.root == nil {
		insertAtRoot(tree, data)
	} else {

		current := tree.root
		for {
			if current.data > data {
				if current.leftChild == nil {
					current.leftChild = createNode(data)
					break
				} else {
					current = current.leftChild
				}
			} else {
				if current.rightChild == nil {
					current.rightChild = createNode(data)
					break
				} else {
					current = current.rightChild
				}
			}
		}
	}
}

//TODO: make insert() on a recursive way

func (tree *BinaryTree) height() int8 {

	var leftHeight, rightHeight int8

	current := tree.root
	for current.leftChild != nil {
		leftHeight++
		current = current.leftChild
	}

	current = tree.root
	for current.rightChild != nil {
		rightHeight++
		current = current.rightChild
	}

	if leftHeight > rightHeight {
		return leftHeight
	} else {
		return rightHeight
	}

}

// TODO: make a InOrderTraversal method

func main() {}
