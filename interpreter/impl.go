package interpreter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/c1rno/algorithm/linked_list"
)

// handle brackets, can't into operations priority
func UnpriorityCalculator(expression string) (float64, error) {
	const valSize = 64
	formatFloat := func(f float64) string {
		r := strconv.FormatFloat(f, 'f', 5, valSize)
		fmt.Printf("%v\n", r)
		return r
	}

	op := linked_list.New()
	val := linked_list.New()

	for _, token := range strings.Split(expression, " ") {
		switch token {
		case "(":
			continue
		case "+", "-", "*", "/":
			op.Push(token)
		case ")":
			rightValStr := val.Pop()
			rightVal, err := strconv.ParseFloat(rightValStr, valSize)
			if err != nil {
				return 0, fmt.Errorf("cant parse %s: %w", rightValStr, err)
			}

			leftValStr := val.Pop()
			leftVal, err := strconv.ParseFloat(leftValStr, valSize)
			if err != nil {
				return 0, fmt.Errorf("cant parse %s: %w", leftValStr, err)
			}

			switch op.Pop() {
			case "+":
				val.Push(formatFloat(leftVal + rightVal))
			case "-":
				val.Push(formatFloat(leftVal - rightVal))
			case "*":
				val.Push(formatFloat(leftVal * rightVal))
			case "/":
				val.Push(formatFloat(leftVal / rightVal))
			}
		default:
			val.Push(token)
		}
	}

	return strconv.ParseFloat(val.Pop(), valSize)
}
