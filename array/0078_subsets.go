package array

func subsets(nums []int) [][]int {
	c, res := []int{}, [][]int{}
	for i := 0; i <= len(nums); i++ {
		generateSubsets(nums, i, 0, c, &res)
	}
	return res
}

func generateSubsets(nums []int, i, start int, c []int, res *[][]int) {
	if len(c) == i {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for j := start; j < len(nums)-(i-len(c))+1; j++ {
		c = append(c, nums[j])
		generateSubsets(nums, i, j+1, c, res)
		c = c[:len(c)-1]
	}
	return
}
