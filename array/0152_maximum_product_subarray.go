package array

func maxProduct(nums []int) int {
	minimum, maximum, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maximum, minimum = minimum, maximum
		}
		maximum = max152(nums[i], maximum*nums[i])
		minimum = min152(nums[i], minimum*nums[i])
		res = max152(res, maximum)
	}
	return res
}

func max152(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min152(x, y int) int {
	if x < y {
		return x
	}
	return y
}
