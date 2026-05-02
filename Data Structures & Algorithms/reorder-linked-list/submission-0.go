/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	arr := []*ListNode{}
	
	// loop list and insert to array
	curr := head
	for curr != nil{
		arr = append(arr, curr)
		curr = curr.Next
	}

	// array index
	l, r := 1, len(arr) - 1

	// reset pointer
	curr = head

	for i:= 0; i < len(arr) - 1; i++ {
		var next int
		if i % 2 == 0 {
			next = r
			r = r - 1
		} else {
			next = l
			l = l + 1
		}

		curr.Next = arr[next]
		curr = arr[next]
	}

	curr.Next = nil
}
