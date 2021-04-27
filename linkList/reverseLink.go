package linkList

func reverseList(head *ListNode) *ListNode {
	if  head == nil || head.Next == nil {
		return head
	}
	p1 := head
	p2 := p1.Next
	for p2 != nil {
		p3 := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}
	head.Next = nil
	return p1
}