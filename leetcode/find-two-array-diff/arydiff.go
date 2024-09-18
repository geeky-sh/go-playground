package findtwoarraydiff

// ref: https://leetcode.com/problems/find-the-difference-of-two-arrays/description/?envType=study-plan-v2&envId=leetcode-75

func findDifference(nums1 []int, nums2 []int) [][]int {
	nmap1 := map[int]int{}
	for _, v := range nums1 {
		_, ok := nmap1[v]
		if !ok {
			nmap1[v] = 0
		}
		nmap1[v] += 1
	}

	nmap2 := map[int]int{}
	for _, v := range nums2 {
		_, ok := nmap2[v]
		if !ok {
			nmap2[v] = 0
		}
		nmap2[v] += 1
	}

	ans1 := []int{}
	c1 := map[int]int{}
	for _, v := range nums1 {
		if _, ok := c1[v]; ok {
			continue
		}
		c1[v] = 1
		_, ok := nmap2[v]
		if !ok {
			ans1 = append(ans1, v)
		}
	}

	ans2 := []int{}
	c2 := map[int]int{}
	for _, v := range nums2 {
		if _, ok := c2[v]; ok {
			continue
		}
		c2[v] = 1
		_, ok := nmap1[v]
		if !ok {
			ans2 = append(ans2, v)
		}
	}

	return [][]int{ans1, ans2}
}
