package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	curr := head
	leftStart, rightStart := &ListNode{}, &ListNode{}
	left, right := leftStart, rightStart

	for curr != nil {
		if curr.Val < x {
			left.Next = curr
			left = left.Next
		} else {
			right.Next = curr
			right = right.Next
		}
		curr = curr.Next
	}

	left.Next = rightStart.Next
	right.Next = nil
	return leftStart.Next
}
