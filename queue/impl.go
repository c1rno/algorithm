package queue

type FIFO struct {
	s []int
}

func New(size int) *FIFO {
	return &FIFO{s: make([]int, 0, size)}
}

func (f *FIFO) Enqueue(i int) {
	f.s = append(f.s, i)
}

func (f *FIFO) Dequeue() (i int) {
	if len(f.s) == 0 {
		return
	}

	i, f.s = f.s[0], f.s[1:]
	return
}
