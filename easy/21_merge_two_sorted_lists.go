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
	// dummy to hold the merged linked list
	dummyNode := new(ListNode)

	// pointer to traverse the two linked lists
	cur := dummyNode

	// loop until either of list become nil
	// to ensure we can compare value of list1 and list2
	for list1 != nil && list2 != nil {
		// if list1 has smaller value
		// add list1 to the next cur pointer
		// and move the list1 pointer forward
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else { // else do the opposite
			cur.Next = list2
			list2 = list2.Next
		}

		// move pointer forward
		cur = cur.Next
	}

	// add the rest of opposite list if no more next
	if list1 == nil {
		cur.Next = list2
	} else {
		cur.Next = list1
	}

	// return the next because it's the merged list
	// dummynode is just a wrapper
	return dummyNode.Next
}

// @lc code=end
