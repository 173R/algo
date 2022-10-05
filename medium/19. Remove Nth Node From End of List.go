func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left := head
	right := head

	for i := 0; i < n; i++ {
		right = right.Next
	}

	if right == nil {
		head = head.Next
	} else {
		for {
			if right.Next == nil {
				break
			} else {
				left = left.Next
				right = right.Next
			}
		}

		left.Next = left.Next.Next
	}

	return head
}