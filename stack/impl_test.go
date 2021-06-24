package stack

import (
	"fmt"
	"testing"

	"github.com/c1rno/algorithm/utils"
)

func create() *LIFO {
	q := New(10)

	q.Push(7)
	q.Push(8)
	q.Push(9)

	return q
}

func TestPop(t *testing.T) {
	q := create()

	utils.Assert("9", q.Pop() == 9)
	utils.Assert("8", q.Pop() == 8)
	utils.Assert("7", q.Pop() == 7)
	utils.Assert("0", q.Pop() == 0)
}

func Example() {
	q := create()

	fmt.Printf("%#v\n", *q)
	// Output:
	// stack.LIFO{s:[]int{7, 8, 9}}
}
