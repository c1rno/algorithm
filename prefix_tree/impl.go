package prefix_tree

import (
	"fmt"
)

// order matter
type list struct {
	s *[]*Node
}

func newList(size int) list {
	s := make([]*Node, 0, size)
	return list{s: &s}
}

func (s list) find(b byte) *Node {
	for _, n := range *s.s {
		if b == n.d {
			return n
		}
	}
	return nil
}

func (s list) index(i int) *Node {
	if i < s.len() {
		slice := *s.s
		return slice[i]
	}
	return nil
}

func (s list) len() int {
	return len(*s.s)
}

func (s list) add(n *Node) {
	*s.s = append(*s.s, n)
}

type listIter struct {
	list
	pos int
}

func newListIter(l list) *listIter {
	return &listIter{list: l, pos: 0}
}

func (i *listIter) next() *Node {
	n := i.list.index(i.pos)
	i.pos++
	return n
}

type Node struct {
	id uint
	// parent
	p *Node
	// children
	c list
	// data
	d byte
}

var count = func() func() uint {
	var c uint
	return func() uint {
		r := c
		c++
		return r
	}
}()

func newNode() *Node {
	const defaultSize = 2
	return &Node{id: count(), c: newList(defaultSize)}
}

// return root-node
func New(words ...string) *Node {
	r := newNode()

	for _, word := range words {
		curr := r
		for _, b := range []byte(word) {
			curr = curr.Add(b)
		}
	}

	return r
}

func (n *Node) Add(b byte) *Node {
	existingChild := n.c.find(b)
	if nil == existingChild {
		newChild := newNode()
		newChild.p = n
		newChild.d = b
		n.c.add(newChild)

		return newChild
	}

	return existingChild
}

func (n *Node) Path(p []byte) *Node {
	if len(p) == 0 {
		return n
	}

	curr := n.c.find(p[0])
	if nil == curr {
		return nil
	}

	return curr.Path(p[1:])
}

func (n *Node) String() string {
	pid := "nil"
	if nil != n.p {
		pid = fmt.Sprintf("%d", n.p.id)
	}
	d := "''"
	if 0 != n.d {
		d = string(n.d)
	}
	return fmt.Sprintf("%s, <%d>, %s, %d", pid, n.id, d, n.c.len())
}

func (n *Node) Extract() []string {
	ch := make([]*Node, 0, 2)
	children(n, &ch)

	words := make([]string, 0, 2)
	for _, child := range ch {
		word := make([]byte, 0, 2)
		arise(child, &word)
		words = append(words, string(word))
	}

	return words
}

func children(n *Node, ch *[]*Node) {
	it := newListIter(n.c)
	for c := it.next(); c != nil; c = it.next() {
		if c.c.len() != 0 {
			children(c, ch)
		} else {
			*ch = append(*ch, c)
		}
	}
}

func arise(n *Node, w *[]byte) {
	if nil == n.p {
		return
	}

	arise(n.p, w)
	*w = append(*w, n.d)
}
