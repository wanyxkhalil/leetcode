package array

import (
	s "github.com/halfrost/LeetCode-Go/structures"
)

func buildTree106(inorder []int, postorder []int) *s.TreeNode {
	inPos := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inPos[inorder[i]] = i
	}
	return buildInPos2TreeDFS(postorder, 0, len(postorder)-1, 0, inPos)
}

func buildInPos2TreeDFS(post []int, postStart, postEnd, inStart int, inPos map[int]int) *s.TreeNode {
	if postStart > postEnd {
		return nil
	}

	root := &s.TreeNode{Val: post[postEnd]}
	rootIdx := inPos[post[postEnd]]
	leftLen := rootIdx - inStart
	root.Left = buildInPos2TreeDFS(post, postStart, postStart+leftLen-1, inStart, inPos)
	root.Right = buildInPos2TreeDFS(post, postStart+leftLen, postEnd-1, rootIdx+1, inPos)
	return root
}
