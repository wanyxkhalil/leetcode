package array

func twoSum167(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] == target {
			return []int{i + 1, j + 1}
		} else if nums[i]+nums[j] < target {
			i++
		} else {
			j--
		}
	}
	return []int{-1, -1}
}
