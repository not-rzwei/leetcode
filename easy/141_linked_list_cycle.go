//go:build ignore
// +build ignore

/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */
package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func hasCycle(head *ListNode) bool {
	// create pointer that moves at different speed
	// slow will move 1x
	// fast will move 2x
	slow, fast := head, head

	// loop until fast finished traversing the list
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// break the cycle if slow and fast meet each other
		// since fast is supposed to finish traversing
		if slow == fast {
			return true
		}
	}

	return false
}

// @lc code=end
