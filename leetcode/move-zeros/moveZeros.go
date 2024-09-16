package movezeros

// https://leetcode.com/problems/move-zeroes/?envType=study-plan-v2&envId=leetcode-75

func moveZeroes(nums []int) {
	zs := -1
	for i, n := range nums {
		if n == 0 {
			if zs == -1 {
				zs = i
			}
		} else {
			if zs != -1 {
				// swap i and zs
				t := nums[zs]
				nums[zs] = nums[i]
				nums[i] = t

				zs += 1
			}
		}
	}
}
