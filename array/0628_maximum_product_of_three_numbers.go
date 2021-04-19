package array

import "sort"

func maximumProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 1
	if len(nums) <= 3 {
		for i := 0; i < len(nums); i++ {
			res = res * nums[i]
		}
		return res
	}
	sort.Ints(nums)
	// if nums[len(nums)-1] <= 0 {
	// 	return 0
	// }

	return max628(nums[0]*nums[1]*nums[len(nums)-1], nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3])
}

func max628(a, b int) int {
	if a > b {
		return a
	}
	return b
}
