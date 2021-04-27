package linkList

func sortList(head *ListNode) *ListNode {
	// 归并排序
	if head == nil || head.Next == nil {
		return head
	}

    mid := middleNode(head)
    rightHead := mid.Next
    mid.Next = nil

	left := sortList(head)
    right := sortList(rightHead)

    return mergeTwoLists(left, right)
}


func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next.Next
	for fast != nil  && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := new(ListNode)
	p := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
			p = p.Next
		} else {
			p.Next = l2
			l2 = l2.Next
			p = p.Next
		}
	}
	if l1 != nil {
		p.Next = l1
	} else {
		p.Next = l2
	}
    return head.Next
}
