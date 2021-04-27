package linkList

type ListNode struct {
	Val int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	res := make([]int, 0)
	// 递归
	var travel func(node *ListNode)
	travel = func(node *ListNode) {
		if node.Next != nil {
			travel(node.Next)
		}
		res = append(res, node.Val)
	}
	travel(head)
	return res
}
