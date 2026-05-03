/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    curr := head
	
	result := &Node{}

	hashMap := make(map[*Node]*Node, 0)

	cursor := result

	for curr != nil {
		cursor.Next = &Node{
			Val: curr.Val,
		}

		// hashMap
		hashMap[curr] = cursor.Next

		// iterate
		cursor = cursor.Next
		curr = curr.Next
	}

	// reset current & cursor to head/result
	curr = head
	cursor = result

	for curr != nil {
		if curr.Random == nil {
			cursor.Next.Random = nil
		} else {
			cursor.Next.Random = hashMap[curr.Random]
		}

		// iterate
		cursor = cursor.Next
		curr = curr.Next
	}

	return result.Next
}
