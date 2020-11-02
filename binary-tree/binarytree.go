package binarytree

import (
	"fmt"
)

// 按照层序遍历生成完全二叉树
func GenerateCompleteBinaryTree(vals []*int) *TreeNode { return nil }

func main() {
	t4 := &TreeNode{Val: 4, Left: nil, Right: nil}
	t5 := &TreeNode{Val: 5, Left: nil, Right: nil}
	t6 := &TreeNode{Val: 6, Left: nil, Right: nil}
	t7 := &TreeNode{Val: 7, Left: nil, Right: nil}
	t2 := &TreeNode{Val: 2, Left: t4, Right: t5}
	t3 := &TreeNode{Val: 3, Left: t6, Right: t7}
	root := &TreeNode{1, t2, t3}

	result := LevelOrder(root)
	for level, result := range result {
		fmt.Printf("level %d result %v\n", level, result)
	}
}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	// 通过上一层的长度确定下一层的元素
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		// 为什么要取length？
		// 记录当前层有多少元素（遍历当前层，再添加下一层）
		l := len(queue)
		for i := 0; i < l; i++ {
			// 出队列
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		result = append(result, list)
	}
	return result
}
