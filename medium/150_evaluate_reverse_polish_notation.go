/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */

package medium

import (
	"strconv"
)

// @lc code=start

type Operator func(x, y int) int
type Stack []int

func (s *Stack) Add(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	last := len(*s) - 1
	popped := (*s)[last]
	*s = (*s)[:last]
	return popped
}

func evalRPN(tokens []string) int {
	// map operator to its function
	// used when we do the calculation
	operators := map[string]Operator{
		"+": func(x, y int) int {
			return x + y
		},
		"-": func(x, y int) int {
			return x - y
		},
		"*": func(x, y int) int {
			return x * y
		},
		"/": func(x, y int) int {
			return x / y
		},
	}

	// create stack to when to add number and to perform operation
	stack := make(Stack, 0)

	// loop through the tokens
	for _, token := range tokens {
		// if token is an operator
		// get the last 2 number from stack
		// and perform operation
		// add back result to stack
		if operator, isOperator := operators[token]; isOperator {
			y, x := stack.Pop(), stack.Pop()
			stack = append(stack, operator(x, y))
			continue
		}

		// add number to stack by convertint the string to int
		number, _ := strconv.Atoi(token)
		stack.Add(number)
	}

	// return the first element because every token has been exhausted
	return stack[0]
}

// @lc code=end
