package utils

import "fmt"

func Assert(msg string, cond bool) {
	if cond {
		return
	}

	panic(fmt.Sprint("Assertion: ", msg))
}
