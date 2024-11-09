/*
 * @lc app=leetcode id=496 lang=golang
 *
 * [496] Next Greater Element I
 */

package easy

// @lc code=start
// stack implementation for ease of use
type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Top() int {
	lastIdx := len(*s) - 1
	return (*s)[lastIdx]
}

func (s *Stack) Pop() int {
	lastIdx := len(*s) - 1
	popped := s.Top()
	*s = (*s)[:lastIdx]
	return popped
}

func nextGreaterElement(nums1 []int, nums2 []int) (result []int) {
	// create a monotonically decreasing stack
	// e.g. [5, 4, 3, 2]
	stack := make(Stack, 0, len(nums2))

	// map number to its next greater number
	greaterMap := make(map[int]int, len(nums1))

	// loop from end to start to get the biggest number
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]

		// enforce monotonic property
		// the stack values have to be decreasing
		// by popping the stack if number on top of stack is lower than current number
		// e.g. [5, 4, 3] to push 6, remove 3 until 5 (pop until empty)
		for len(stack) > 0 && stack.Top() <= num {
			stack.Pop()
		}

		// if empty, current number has no greater element
		// because it's bigger than any number after it
		if len(stack) == 0 {
			greaterMap[num] = -1
		} else {
			// else, we have greater number after it
			greaterMap[num] = stack.Top()
		}

		// add the number to monotonic stack
		stack.Push(num)
	}

	// get each greater number from map and add it to the result
	for _, num := range nums1 {
		result = append(result, greaterMap[num])
	}

	return
}

// @lc code=end
