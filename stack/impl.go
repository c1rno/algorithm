package stack

type LIFO struct {
	s []int
}

func New(size int) *LIFO {
	return &LIFO{s: make([]int, 0, size)}
}

func (f *LIFO) Push(i int) {
	f.s = append(f.s, i)
}

func (f *LIFO) Pop() (i int) {
	if len(f.s) == 0 {
		return
	}

	f.s, i = f.s[:len(f.s)-1], f.s[len(f.s)-1]
	return
}
