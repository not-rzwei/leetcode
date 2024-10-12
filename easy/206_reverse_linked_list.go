//go:build ignore
// +build ignore

/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */
package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func reverseList(head *ListNode) *ListNode {
	// pointer to current linked list position
	cur := head

	// pointer to the previous node of linked list
	// we initiate with empty node because first previous is empty
	var prev *ListNode

	// loop until we traverse the linked list
	// basically when the current next is empty
	for cur != nil {
		// switch the next pointer to prev
		// set prev to current
		// and move the current pointer forward
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

// @lc code=end
