package findhighestaltitude

func largestAltitude(gain []int) int {
	alts := []int{0}
	res := 0
	for i, gn := range gain {

		nv := alts[i] + gn
		alts = append(alts, nv)

		if nv > res {
			res = nv
		}
	}
	return res
}
