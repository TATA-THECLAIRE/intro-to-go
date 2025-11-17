package main

import "fmt"

// TreeNode represents a single node in the binary tree
type TreeNode struct {
	Value int       // The data stored in this node
	Left  *TreeNode // Pointer to the left child
	Right *TreeNode // Pointer to the right child
}

// IsBalanced checks if the entire tree is balanced
// A tree is balanced if every node's left and right subtrees differ in height by at most 1(0 or 1)
func IsBalanced(root *TreeNode) bool {
	// If the tree is empty, it's balanced
	if root == nil {
		return true
	}
	
	// Checking if this tree is balanced by getting the height
	// If height is -1, the tree is not balanced
	height := checkHeight(root)
	return height != -1
}

// checkHeight calculates the height of a tree
// Returns -1 if the tree is NOT balanced
// Returns the actual height if the tree IS balanced
func checkHeight(node *TreeNode) int {
	// Base case: an empty tree has height 0
	if node == nil {
		return 0
	}
	
	// Check the left subtree
	leftHeight := checkHeight(node.Left)
	if leftHeight == -1 {
		// Left subtree is not balanced, so this tree isn't either
		return -1
	}
	
	// Check the right subtree
	rightHeight := checkHeight(node.Right)
	if rightHeight == -1 {
		// Right subtree is not balanced, so this tree isn't either
		return -1
	}
	
	// Calculate the difference in heights
	difference := leftHeight - rightHeight
	if difference < 0 {
		difference = -difference // Make it positive
	}
	
	// If the difference is more than 1, tree is not balanced
	if difference > 1 {
		return -1
	}
	
	// Tree is balanced! Return the height
	// Height = 1 (current node) + max(left height, right height)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// Helper function to create a new node
func NewNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func main() {
	// Example 1: Balanced tree
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	balanced := NewNode(1)
	balanced.Left = NewNode(2)
	balanced.Right = NewNode(3)
	balanced.Left.Left = NewNode(4)
	balanced.Left.Right = NewNode(5)
	
	fmt.Println("Example 1 - Balanced tree:")
	fmt.Printf("Is balanced? %v\n\n", IsBalanced(balanced))
	
	// Example 2: Unbalanced tree
	//       1
	//      /
	//     2
	//    /
	//   3
	unbalanced := NewNode(1)
	unbalanced.Left = NewNode(2)
	unbalanced.Left.Left = NewNode(3)
	
	fmt.Println("Example 2 - Unbalanced tree:")
	fmt.Printf("Is balanced? %v\n\n", IsBalanced(unbalanced))
	
	// Example 3: Empty tree
	fmt.Println("Example 3 - Empty tree:")
	fmt.Printf("Is balanced? %v\n", IsBalanced(nil))
}