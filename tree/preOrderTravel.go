package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	var st1 []*TreeNode
	var st2 []*TreeNode
	st1 = append(st1, root)
	for level := 0; len(st1) > 0 || len(st2) > 0; level++ {
		res = append(res, []int{})
		n := len(st1)
		for i := 0; i < n; i++ {
			node := st1[len(st1)-1]
			st1 = st1[:len(st1)-1]
			res[level] = append(res[level], node.Val)
			if node.Left != nil {
				st2 = append(st2, node.Left)
			}
			if node.Right != nil {
				st2 = append(st2, node.Right)
			}

		}
		m := len(st2)
		for i := 0; i < m; i++ {
			// pop
			if i == 0 {
				level ++
				res = append(res, []int{})
			}
			node := st2[len(st2)-1]
			st2 = st2[:len(st2)-1]

			res[level] = append(res[level], node.Val)
			if node.Right != nil {
				st1 = append(st1, node.Right)
			}
			if node.Left != nil {
				st1 = append(st1, node.Left)
			}
		}

	}
	return res
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}

	for level := 0; 0 < len(queue); level++ {
		res = append(res, []int{})
		next := []*TreeNode{}
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			if node != nil {
				res[level] = append(res[level], node.Val)
				if node.Left != nil {
					next = append(next, node.Left)
				}
				if node.Right != nil {
					next = append(next, node.Right)
				}
			}
		}
		queue = next
	}
	return res
}
