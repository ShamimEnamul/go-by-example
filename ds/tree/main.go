package main

import (
	"fmt"
	"math"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func getNewNode(val int) *Node {
	return &Node{
		val: val,
	}
}

func insert(root *Node, val int) *Node {
	if root == nil {
		root = getNewNode(val)
	} else if val < root.val {
		root.left = insert(root.left, val)
	} else {
		root.right = insert(root.right, val)
	}

	return root
}

func (root *Node) insert(val int) *Node {
	if root == nil {
		root = getNewNode(val)
	} else if val < root.val {
		root.left = root.left.insert(val)
	} else {
		root.right = root.right.insert(val)
	}

	return root
}

func findMin(root *Node) int {
	if root == nil {
		return -1
	}
	for root.left != nil {
		root = root.left
	}
	return root.val
}

func findMinRec(root *Node) int {
	if root == nil {
		return -1
	}
	for root.left == nil {
		return root.val
	}
	return findMinRec(root.left)
}

func max(left int, right int) int {
	if right > left {
		return right
	}
	return left
}
func findHeight(root *Node) int {

	if root == nil {
		return -1
	}
	leftVal := findHeight(root.left)
	rightVal := findHeight(root.right)
	return 1 + max(leftVal, rightVal)
}

func show(root *Node) {
	if root == nil {
		return
	}

	show(root.left)
	fmt.Println(root.val)
	show(root.right)
}

// Level Order Traverse
func levelOrder(root *Node) {
	if root == nil {
		return
	}

	q := Queue{}
	q.push(root)

	for root != nil {
		curr := q.front()
		if curr == nil {
			break
		}
		fmt.Println(curr.val)
		q.push(curr.left)
		q.push(curr.right)
		q.pop()
	}
}

func isBstUtils(root *Node, minVal int, maxVal int) bool {

	if root == nil {
		return true
	}

	if root.val >= minVal &&
		root.val <= maxVal &&
		isBstUtils(root.left, minVal, root.val) &&
		isBstUtils(root.right, root.val, maxVal) {
		return true
	}

	return false
}
func isBst(root *Node) bool {

	return isBstUtils(root, math.MinInt, math.MaxInt)
}

func findVal(root *Node, data int) *Node {
	for root != nil {
		if root.val == data {
			return root
		} else if root.val > data {
			root = root.left
		} else {
			root = root.right
		}
	}
	return nil
}

func findMinNode(root *Node) *Node {
	if root == nil {
		return nil
	}
	for root.left != nil {
		root = root.left
	}
	return root
}

// delete a node
func deleteNode(root *Node, data int) *Node {
	// travel to the node
	if root == nil {
		return nil
	} else if root.val < data {
		root.right = deleteNode(root.right, data)
	} else if root.val > data {
		root.left = deleteNode(root.left, data)
	} else {
		if root.left == nil && root.right == nil {
			root = nil
		} else if root.left == nil {
			root = root.right

		} else if root.right == nil {
			root = root.left
		} else {
			temp := findMinNode(root.right)
			root.val = temp.val
			root.right = deleteNode(root.right, temp.val)
		}

	}

	return root
	// get that max of root.right or min from the left right and replace with it
}

func main() {
	fmt.Println("::: TREE :::")
	root := &Node{val: 15}
	root = insert(root, 10)
	root = insert(root, 9)
	root = insert(root, 11)
	root = insert(root, 6)
	root = insert(root, 8)
	//show(root)
	//fmt.Println("min: REC : ", findMinRec(root))
	//
	fmt.Println(isBst(root))
	//levelOrder(root)
}
