/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    

	// find the middle node right by using fast and slow pointer
	// the middle node right is slow
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse the second reorderList
	// fmt.Printf("fast: %v, slow: %v", fast, slow)
	// reverse list
	secondList := slow.Next
	slow.Next = nil
	var prevNode *ListNode
	for secondList != nil {
		// fmt.Printf("secondList val: %v\n", secondList.Val)
		tmp := secondList.Next
		secondList.Next = prevNode
		prevNode = secondList
		secondList = tmp
	}

	firstList := head
	secondList = prevNode

	for secondList != nil {
		tmp1, tmp2 := firstList.Next, secondList.Next
		firstList.Next = secondList
		secondList.Next = tmp1
		firstList, secondList = tmp1, tmp2 // to run loop
	}

	for curr := head; curr != nil; curr = curr.Next {
		fmt.Printf("head val: %v\n", curr.Val)
	}
	
}
