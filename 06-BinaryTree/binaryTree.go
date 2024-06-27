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

// func heightNonRecursive(root *Node) int {
// 	if root == nil {
// 		return 0
// 	}

// 	queue := []*Node{root}
// 	height := 0

// 	for len(queue) > 0 {
// 		size := len(queue)
// 		for i := 0; i < size; i++ {
// 			node := queue[0]
// 			queue = queue[1:]

// 			if node.leftChild != nil {
// 				queue = append(queue, node.leftChild)
// 			}
// 			if node.rightChild != nil {
// 				queue = append(queue, node.rightChild)
// 			}
// 		}
// 		height++
// 	}
// 	return height
// }

func height(node *Node) int8 {
	if node == nil {
		return 0
	} else {

		var lSide, rDepth = height(node.leftChild), height(node.rightChild)

		if lSide > rDepth {
			return (lSide + 1)
		} else {
			return (rDepth + 1)
		}
	}
}

func (tree *BinaryTree) inorderTraversal() []int8 {
	var output = []int8{}
	var stack = []*Node{}
	current := tree.root

	for current != nil || len(stack) > 0 {

		//put all leftNodes in the stack
		for current != nil {

			stack = append(stack, current)
			current = current.leftChild
		}

		//return current to the fatherNode(previus Current), becuse there is no more leftNode in the subtree
		current = stack[len(stack)-1]

		//remove last index from stack
		stack = stack[:len(stack)-1]

		output = append(output, current.data)

		//visit the rightNode
		current = current.rightChild

	}
	return output

}

func (tree *BinaryTree) preorderTraversal() []int8 {
	var output = []int8{}
	var stack = []*Node{}
	current := tree.root

	for current != nil || len(stack) > 0 {

		for current != nil {
			stack = append(stack, current)
			output = append(output, current.data)

			current = current.leftChild
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		current = current.rightChild
	}

	return output

}

func (tree *BinaryTree) postorderTraversal() []int8 {
	var output []int8
	var stack []*Node

	var lastVisited *Node

	current := tree.root

	for current != nil || len(stack) > 0 {

		for current != nil {
			stack = append(stack, current)
			current = current.leftChild
		}

		current = stack[len(stack)-1]

		if current.rightChild == nil || current.rightChild == lastVisited {

			stack = stack[:len(stack)-1]
			output = append(output, current.data)
			lastVisited = current
			current = nil
		} else {
			current = current.rightChild
		}
	}

	return output
}

// TODO: make a Level Order Traversal method
func (tree *BinaryTree) leverOrderTraversal() []int8 {
	var output = []int8{}
	var queue = []*Node{}

	queue = append(queue, tree.root)

	for len(queue) > 0 {

		current := queue[0]

		output = append(output, current.data)

		if current.leftChild != nil {
			queue = append(queue, current.leftChild)
		}

		if current.rightChild != nil {
			queue = append(queue, current.rightChild)
		}

		queue = queue[1:]

	}

	return output

}

func main() {}
