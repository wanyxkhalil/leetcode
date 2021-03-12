package array

import (
	s "github.com/halfrost/LeetCode-Go/structures"
)

func buildTree(preorder []int, inorder []int) *s.TreeNode {
	inPos := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inPos[inorder[i]] = i
	}
	return buildPreIn2TreeDFS(preorder, 0, len(preorder)-1, 0, inPos)
}

func buildPreIn2TreeDFS(pre []int, preStart, preEnd, inStart int, inPos map[int]int) *s.TreeNode {
	if preStart > preEnd {
		return nil
	}
	root := &s.TreeNode{Val: pre[preStart]}
	rootIdx := inPos[pre[preStart]]
	leftLen := rootIdx - inStart
	root.Left = buildPreIn2TreeDFS(pre, preStart+1, preStart+leftLen, inStart, inPos)
	root.Right = buildPreIn2TreeDFS(pre, preStart+leftLen+1, preEnd, rootIdx+1, inPos)
	return root
}
