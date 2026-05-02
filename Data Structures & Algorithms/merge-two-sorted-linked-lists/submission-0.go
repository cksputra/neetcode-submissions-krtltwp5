/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	head1, head2 := list1, list2

	val := []int{}

	for head1 != nil{
		val = append(val, head1.Val)
		head1 = head1.Next
	}

	for head2 != nil{
		val = append(val, head2.Val)
		head2 = head2.Next
	}

	sort.Ints(val)


	result := &ListNode{}
	current := result
	for _, v := range val{
		next:= &ListNode{
			Val: v,
		}

		current.Next = next
		current = next	
	}

	return result.Next
}
