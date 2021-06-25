package requisition

func Fib(n int) []int {
	r := make([]int, 0, n)

	for i := 0; i < n; i++ {
		r = append(r, fib(i, r))
	}

	return r
}

func fib(n int, reuse []int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return reuse[n-2] + reuse[n-1]
}

func Fac(n int) []int {
	r := make([]int, 0, n)

	for i := 0; i < n; i++ {
		r = append(r, fac(i))
	}

	return r
}

func fac(n int) int {
	if n == 0 {
		return 1
	}
	return n * fac(n-1)
}
