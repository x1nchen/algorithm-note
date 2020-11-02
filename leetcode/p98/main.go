package main

import (
	"fmt"

	"algorithm/binary-tree"
)

func main() {
	t2 := &binarytree.TreeNode{Val: 1, Left: nil, Right: nil}
	t3 := &binarytree.TreeNode{Val: 3, Left: nil, Right: nil}
	root := &binarytree.TreeNode{Val: 2, Left: t2, Right: t3}
	ok := isValidBST(root)
	fmt.Println(ok)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *binarytree.TreeNode) bool {
	if root ==  nil {
		return true
	}

	out := traverse(root)
	if len(out) == 1 {
		return true
	}

	last := out[0]
	for _, i := range out {
		if last > i {
			return false
		}
		last = i
	}
	return true
}

func traverse(root *binarytree.TreeNode) []int {
	var out []int
	if root == nil {
		return out
	}
	out = append(out, traverse(root.Left)...)
	out = append(out, root.Val)
	out = append(out, traverse(root.Right)...)
	return out
}
