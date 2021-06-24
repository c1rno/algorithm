package set

type Set map[string]struct{}

func (s Set) Add(v string)               { s[v] = struct{}{} }
func (s Set) Contain(v string) (ok bool) { _, ok = s[v]; return }
func (s Set) Del(v string)               { delete(s, v) }
