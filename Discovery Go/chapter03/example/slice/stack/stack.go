// Package stack provides basic implementation with stack data structure
package stack

import (
	"strconv"
	"strings"
)

// Eval returns the evaluation result of the given expr.
// The expression can have +, -, *, /, (, ) operators and
// decimal integers. Operators and operands should be
// space delimited.
func Eval(expr string) int {
	var ops []string
	var nums []int

	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}

	reduce := func(higher string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]

			// 1. check priority
			if strings.Index(higher, op) < 0 {
				// operator not in the list
				return
			}

			// 2. check parenthesis
			ops = ops[:len(ops)-1]
			if op == "(" {
				// delete right parenthesis
				return
			}

			// 3. calculate
			b, a := pop(), pop()
			switch op {
			case "+":
				nums = append(nums, a+b)
			case "-":
				nums = append(nums, a-b)
			case "*":
				nums = append(nums, a*b)
			case "/":
				nums = append(nums, a/b)
			}
		}
	}

	for _, token := range strings.Split(expr, " ") {
		switch token {
		case "(":
			ops = append(ops, token)
		case "+", "-":
			reduce("+-*/")
			ops = append(ops, token)
		case "*", "/":
			reduce("*/")
			ops = append(ops, token)
		case ")":
			reduce("+-*/(")
		default:
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}
	reduce("+-*/")
	return nums[0]
}
