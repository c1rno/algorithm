package requisition

import "fmt"

func ExampleFib() {
	r := Fib(10)
	fmt.Println(r)

	// Output:
	// [0 1 1 2 3 5 8 13 21 34]
}

func ExampleFac() {
	r := Fac(10)
	fmt.Println(r)

	// Output:
	// [1 1 2 6 24 120 720 5040 40320 362880]
}
