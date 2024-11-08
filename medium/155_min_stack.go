/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

package medium

// @lc code=start
type MinStack struct {
	stack []int // store the stack number
	min   []int // track the minimum number
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   make([]int, 0),
	}
}

func (s *MinStack) Push(val int) {
	// add number to the stack
	s.stack = append(s.stack, val)

	// add number to min stack to track the lower numbers
	// only add if we don't have lower numbers
	// or current number is lower than last lower numbers
	if len(s.min) == 0 || val <= s.GetMin() {
		s.min = append(s.min, val)
	}
}

func (s *MinStack) Pop() {
	// pop the latest number from stack
	popped := s.Top()
	s.stack = s.stack[:len(s.stack)-1]

	// and remove the lower number if it's the same as popped number
	if popped == s.GetMin() {
		s.min = s.min[:len(s.min)-1]
	}
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.min[len(s.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
