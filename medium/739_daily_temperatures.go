//go:build forward
// +build forward

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

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	monostack := make(Stack, 0)

	// traverse through temperatures
	for i := 0; i < len(temperatures); i++ {
		// pop the stack if stack is not empty and current temperature is higher
		// than what we have on the stack
		// and calculate the distance for the temperature on the stack
		// e.g. [70,69,71] cur 70 stack [70]
		// 		           cur 69 stack [70, 69]
		//		           cur 71 pop stack until condition is satisfied
		//                 cur 71 stack [70] distance 69 = index 71-69 = 1
		//                 cur 71 stack [] distance 70 = index 71-70 = 2
		for !monostack.IsEmpty() && temperatures[i] > temperatures[monostack.Top()] {
			popped := monostack.Pop()
			result[popped] = i - popped
		}

		// add temperature index to the stack
		monostack.Push(i)
	}

	return result
}

// @lc code=end
