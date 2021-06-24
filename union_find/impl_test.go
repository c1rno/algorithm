package union_find

import (
	"fmt"
	"strings"
	"testing"

	"github.com/c1rno/algorithm/utils"
)

func create() *UF {
	const defaultSize = 10

	u := New(defaultSize)
	u.Union(1, 2)
	u.Union(2, 3)
	u.Union(3, 4)

	u.Union(0, 9)

	return u
}

func TestFind(_ *testing.T) {
	u := create()

	utils.Assert("1, 2 connected", u.Connected(1, 2))
	utils.Assert("2, 3 connected", u.Connected(2, 3))
	utils.Assert("3, 4 connected", u.Connected(3, 4))
	utils.Assert("1, 4 connected", u.Connected(1, 4))
	utils.Assert("1, 3 connected", u.Connected(1, 3))

	utils.Assert("0, 9 connected", u.Connected(0, 9))
	utils.Assert("0, 1 not connected", !u.Connected(0, 1))
	utils.Assert("0, 2 not connected", !u.Connected(0, 2))
	utils.Assert("0, 3 not connected", !u.Connected(0, 3))
}

func Example() {
	u := create()

	pairs := make([]string, 0, len(u.ids))

	for i, j := range u.ids {
		pairs = append(pairs, fmt.Sprintf("%d, %d, %d", i, j, u.root(i)))
	}

	fmt.Printf("%v\n%s\n", u.ids, strings.Join(pairs, "\n"))
	// Output:
	// [9 2 3 4 4 5 6 7 8 9]
	// 0, 9, 9
	// 1, 2, 4
	// 2, 3, 4
	// 3, 4, 4
	// 4, 4, 4
	// 5, 5, 5
	// 6, 6, 6
	// 7, 7, 7
	// 8, 8, 8
	// 9, 9, 9
}
