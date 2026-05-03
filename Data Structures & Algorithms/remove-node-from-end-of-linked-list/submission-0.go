/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := head

	for i:=0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil{
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

    return dummy.Next
}


// 1 2 3 4 n = 3

// 1
// 1

// 1 1 1
// 2


// 1 2 3 4 5 6 7 8 9   n = 3

// 4 = first
// 1

// 5 6 7 8 9
// 2 3 4 5 6

