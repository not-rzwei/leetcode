//go:build backward
// +build backward

/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */

package medium

// @lc code=start
// stack implementation for cleaner interface
type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	lastIdx := len(*s) - 1
	popped := (*s)[lastIdx]
	*s = (*s)[:lastIdx]
	return popped
}

func (s *Stack) Top() int {
	lastIdx := len(*s) - 1
	return (*s)[lastIdx]
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	monostack := make(Stack, 0, n)

	// traverse temperatures backward
	for i := n - 1; i >= 0; i-- {
		// pop the stack until monodecreasing stack constraint is satisfied
		// e.g. [70,72,71] cur 71 stack [] empty stack
		//                 cur 71 stack [] add 71 to stack
		// 		           cur 72 stack [71] 72 is higher than 71, pop 71. add
		// 		           cur 72 stack [] add 72 to stack
		//		           cur 70 stack [72] 72 is not higher than 70, skip
		for len(monostack) > 0 && temperatures[i] >= temperatures[monostack.Top()] {
			monostack.Pop()
		}

		// if we have something on the stack, top stack is our higher temperature
		// and assign the distance of that top stack with current temp
		// e.g. [70,72,71] cur 71 stack [] empty stack skip
		// 		           cur 72 stack [] empty stack skip
		//		           cur 70 stack [72] distance idx 72-70 = 1
		if len(monostack) > 0 {
			result[i] = monostack.Top() - i
		}

		// add current temperature index to the stack
		monostack.Push(i)
	}

	return result
}

// @lc code=end
