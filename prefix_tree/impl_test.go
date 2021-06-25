package prefix_tree

import (
	"fmt"
	"testing"

	"github.com/c1rno/algorithm/utils"
)

func create() *Node {
	r := New(
		"word-1",
		"word-2",
		"banana",
		"word-3",
		"apple",
		"что-то длинное и непонятное, с пробелами, на русском",
	)

	return r
}

func ExampleNode_Extract() {
	r := create()

	for _, word := range r.Extract() {
		fmt.Println(word)
	}
	// Output:
	// word-1
	// word-2
	// word-3
	// banana
	// apple
	// что-то длинное и непонятное, с пробелами, на русском
}

func ExampleNode_Path() {
	r := create()
	n := r.Path([]byte("wor"))

	for _, word := range n.Extract() {
		fmt.Println(word)
	}
	// Output:
	// word-1
	// word-2
	// word-3
}

func TestNotFind(_ *testing.T) {
	r := create()
	n := r.Path([]byte("won"))

	utils.Assert("no suggestions for 'won'", nil == n)
}
