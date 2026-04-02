/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
	current := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {

		// vì l1 hoặc l2 có độ dài node khác nhau nên cần check như sau
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		// tính tổng
		sum := val1 + val2 + carry
		// nhớ vào carry để công ở node tiếp
		carry = sum / 10

		// next nodes
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	// for dummy != nil {
	// 	fmt.Printf("%v - ", dummy.Val)
	// 	dummy = dummy.Next
	// }
	// fmt.Println()
	return dummy.Next
}
