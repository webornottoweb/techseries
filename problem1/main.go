// check if provided tree is a valid binary search tree
package main

import (
	"fmt"
	"math"
)

// Node struct to be passed through
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func isValidBSTHelper(node *Node, low int, high int) bool {
	if node == nil {
		return true
	}

	val := node.Val

	if (val > low && val < high) && isValidBSTHelper(node.Left, low, val) && isValidBSTHelper(node.Right, val, high) {
		return true
	}

	return false
}

func isValidBST(node *Node) bool {
	return isValidBSTHelper(node, math.MinInt64, math.MaxInt64)
}

func main() {
	/* Valid tree
		5
	  /   \
	4       7
	*/
	node1 := Node{Val: 5}
	node1.Left = &Node{Val: 4}
	node1.Right = &Node{Val: 7}

	fmt.Println(isValidBST(&node1))

	/* Invalid tree

		5
	  /   \
	4       7
		  /	  \
		2		9
	*/
	node1.Right.Right = &Node{Val: 9}
	node1.Right.Left = &Node{Val: 2}

	fmt.Println(isValidBST(&node1))
}
