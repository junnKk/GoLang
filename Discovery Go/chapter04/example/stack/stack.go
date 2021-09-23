// Package stack provides basic implementation with stack data structure
package stack

import (
	"fmt"
	"strconv"
	"strings"
)

type BinOp func(int, int) int

// Eval returns the evaluation result of the given expr.
// The expression can have +, -, *, /, (, ) operators and
// decimal integers. Operators and operands should be
// space delimited.
// [Refactor - can handle each operator's function]
func Eval(opMap map[string]BinOp, prec PrecMap, expr string) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Eval.", r)
		}
	}()

	ops := []string{"("}
	var nums []int

	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}

	reduce := func(nextOp string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if _, higher := prec[nextOp][op]; nextOp != ")" && !higher {
				// 더 낮은 순위 연산자이므로 여기서 계산 종료
				return
			}
			ops = ops[:len(ops)-1]
			if op == "(" {
				return
			}
			b, a := pop(), pop()
			if f := opMap[op]; f != nil {
				nums = append(nums, f(a, b))
			}
		}
	}

	for _, token := range strings.Fields(expr) {
		if token == "errorCase" {
			panic("Error occurred")
		}
		if token == "(" {
			ops = append(ops, token)
		} else if _, ok := prec[token]; ok {
			reduce(token)
			ops = append(ops, token)
		} else if token == ")" {
			reduce(token)
		} else {
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}

	reduce(")")
	return nums[0]
}

// String set type
type StrSet map[string]struct{}

// Returns a new StrSet
func NewStrSet(strs ...string) StrSet {
	m := StrSet{}
	for _, str := range strs {
		m[str] = struct{}{}
	}
	return m
}

// Map keyed by operator to set of higher precedence operators
type PrecMap map[string]StrSet

func NewEvaluator(opMap map[string]BinOp, prec PrecMap) func(expr string) int {
	return func(expr string) int {
		return Eval(opMap, prec, expr)
	}
}
