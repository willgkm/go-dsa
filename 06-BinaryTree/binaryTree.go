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

func (tree *BinaryTree) insertAtRoot(data int8) {
	newNode := createNode(data)
	tree.root = newNode
}

func (tree *BinaryTree) insert(data int8) {
	if tree.root == nil {
		tree.insertAtRoot(data)
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

func main() {

}
