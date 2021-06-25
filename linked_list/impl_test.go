package linked_list

import (
	"fmt"
	"testing"

	"github.com/c1rno/algorithm/utils"
)

func create() *List {
	l := New()

	l.Push("1")
	l.Push("2")
	l.Push("3")

	l.Queue("-1")
	l.Queue("-2")

	return l
}

func TestPop(t *testing.T) {
	l := create()

	utils.Assert("pop 3", l.Pop() == "3")
	utils.Assert("pop 2", l.Pop() == "2")
	utils.Assert("pop 1", l.Pop() == "1")
	utils.Assert("pop -1", l.Pop() == "-1")
	utils.Assert("pop -2", l.Pop() == "-2")
	utils.Assert("pop 0", l.Pop() == "")

	utils.Assert("head empty", l.h == nil)
	utils.Assert("tail empty", l.t == nil)
}

func TestDequeue(t *testing.T) {
	l := create()

	utils.Assert("dequeue -2", l.Dequeue() == "-2")
	utils.Assert("dequeue -1", l.Dequeue() == "-1")
	utils.Assert("dequeue 1", l.Dequeue() == "1")
	utils.Assert("dequeue 2", l.Dequeue() == "2")
	utils.Assert("dequeue 3", l.Dequeue() == "3")
	utils.Assert("dequeue 0", l.Dequeue() == "")

	utils.Assert("head empty", l.h == nil)
	utils.Assert("tail empty", l.t == nil)
}

func ExampleList_Push() {
	l := create()

	fmt.Printf("head: %s\n", formatNode(l.h))

	it, count := &RTLIterator{l.h}, 0
	for curr := it.Next(); curr != nil; curr = it.Next() {
		fmt.Printf("%d:    %s\n", count, formatNode(curr))
		count++
	}

	fmt.Printf("tail: %s\n", formatNode(l.t))

	// Output:
	// head: <nil>, -2, -1
	// 0:    <nil>, -2, -1
	// 1:    -2, -1, 1
	// 2:    -1, 1, 2
	// 3:    1, 2, 3
	// 4:    2, 3, <nil>
	// tail: 2, 3, <nil>
}

func formatNode(n *Node) string {
	const nilStr = "<nil>"

	if n == nil {
		return nilStr
	}

	l := nilStr
	if n.l != nil {
		l = n.l.data
	}

	c := n.data

	r := nilStr
	if n.r != nil {
		r = n.r.data
	}

	return fmt.Sprintf("%s, %s, %s", l, c, r)
}
