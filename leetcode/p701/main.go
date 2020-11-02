// https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/submissions/
package main

import (
	binarytree "algorithm/binary-tree"
)

func main() {

}

func insertIntoBST(root *binarytree.TreeNode, val int) *binarytree.TreeNode {
	if root == nil {
		root = &binarytree.TreeNode{Val: val}
		return root
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}


