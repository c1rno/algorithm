package queue

import (
	"fmt"
	"testing"

	"github.com/c1rno/algorithm/utils"
)

func create() *FIFO {
	q := New(10)

	q.Enqueue(7)
	q.Enqueue(8)
	q.Enqueue(9)

	return q
}

func TestDequeue(t *testing.T) {
	q := create()

	utils.Assert("7", q.Dequeue() == 7)
	utils.Assert("8", q.Dequeue() == 8)
	utils.Assert("9", q.Dequeue() == 9)
	utils.Assert("0", q.Dequeue() == 0)
}

func Example() {
	q := create()

	fmt.Printf("%#v\n", *q)
	// Output:
	// queue.FIFO{s:[]int{7, 8, 9}}
}
