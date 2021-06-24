// https://d3c33hcgiwev3.cloudfront.net/_b65e7611894ba175de27bd14793f894a_15UnionFind.pdf?Expires=1624665600&Signature=EyKNqWunYMK5sS4R7uh6n3ZY74NgDZ3J70z9bBYptwWE-c5~pG0~bLZ8NV1n~6Rk935O4BMsyBr0jbL3Xqf2sAFKY9JpuWx2U2942NW5rIvdTheLRhYSJ8iIsxEeskCQNhGLgTwhn6~v1dalpeG~FY-AZshoQIc58Iogr88q4gs_&Key-Pair-Id=APKAJLTNE6QMUY6HBC5A

package union_find

type UF struct {
	ids []int
}

// New initialize by (0, n-1) points
func New(n int) *UF {
	ids := make([]int, n)

	for i := range ids {
		ids[i] = i
	}

	return &UF{ids: ids}
}

// Union add connections between p and q
func (u *UF) Union(p, q int) {
	i := u.root(p)
	j := u.root(q)

	u.ids[i] = j
}

// Connected are they connected
func (u *UF) Connected(p, q int) bool {
	return u.root(p) == u.root(q)
}

func (u *UF) root(i int) int {
	for i != u.ids[i] {
		i = u.ids[i]
	}

	return i
}
