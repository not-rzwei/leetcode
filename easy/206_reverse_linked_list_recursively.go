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
	// base case
	// if no more next or no head, return head
	// this will be the head of reversed linked list
	if head == nil || head.Next == nil {
		return head
	}

	// recursive case
	// ask who will be the new head recursively
	// until it hit the base case
	// which return the new hewad
	reversedHead := reverseList(head.Next)

	// set the next.Next to current
	// which basically flips the direction
	head.Next.Next = head

	// and cut next link to break the cycle
	head.Next = nil

	return reversedHead
}

// @lc code=end
