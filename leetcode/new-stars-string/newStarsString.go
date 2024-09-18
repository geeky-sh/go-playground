package newstarsstring

// ref: https://leetcode.com/problems/removing-stars-from-a-string/?envType=study-plan-v2&envId=leetcode-75

type St struct {
	lst []rune
}

func (r *St) push(e rune) {
	r.lst = append(r.lst, e)
}

func (r *St) pop() rune {
	e := r.lst[len(r.lst)-1]
	r.lst = r.lst[:(len(r.lst) - 1)]
	return e
}

func NewSt() *St {
	return &St{[]rune{}}
}

func removeStars(s string) string {
	st := NewSt()
	for _, r := range s {
		if r == rune('*') {
			st.pop()
		} else {
			st.push(r)
		}
	}

	return string(st.lst)
}
