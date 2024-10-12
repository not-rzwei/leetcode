//go:build ignore
// +build ignore

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// base case
	// return opposite list if empty
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	// recursive call
	if list1.Val < list2.Val {
		// if list1 value is lower
		// calculate the next node of list1
		list1.Next = mergeTwoLists(list1.Next, list2)

		// return list1 because it's the head
		return list1
	} else {
		// otherwise do the opposite
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

// @lc code=end
