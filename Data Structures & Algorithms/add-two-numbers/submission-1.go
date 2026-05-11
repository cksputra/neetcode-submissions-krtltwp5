/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil || l2 == nil {
		return nil
	}

	head1 := l1
	head2 := l2

	result := &ListNode{}
	rHead := result
	remainder := false

	for {
		if head1 == nil && head2 == nil {
			break
		}

		


		val1, val2 := 0,0

		if head1 != nil {
			val1 = head1.Val
		}

		if head2 != nil {
			val2 = head2.Val
		}

		sum := val1 + val2

		if remainder{
			sum+=1
			remainder = false
		}


		if sum >= 10 {
			sum = sum % 10
			remainder = true
		}

		temp := &ListNode{
			Val: sum,
		}

		rHead.Next = temp
		rHead = rHead.Next

		// move head forward
		if head1 != nil {
			head1 = head1.Next
		}

		if head2 != nil {
			head2 = head2.Next
		}
	}

	if remainder {
		temp := &ListNode{
			Val: 1,
		}
		rHead.Next = temp
	}
    
	return result.Next
}
