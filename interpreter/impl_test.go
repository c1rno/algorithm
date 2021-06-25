package interpreter

import (
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		expression string
		result     float64
		error      error
	}{
		{
			expression: "( 1 + 1 )",
			result:     2,
		},
		{
			expression: "( 1 + ( 1 * 2 ) )",
			result:     3,
		},
		{
			expression: "( ( 3 / 6 ) + ( 1 * 2 ) )",
			result:     2.5,
		},
	}

	for i, c := range cases {
		c := c

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r, err := UnpriorityCalculator(c.expression)

			if r != c.result {
				t.Errorf("unexpected result: %v != %v", r, c.result)
			}

			if err != c.error {
				t.Errorf("unexpected error: %v != %v", err, c.error)
			}
		})
	}
}
