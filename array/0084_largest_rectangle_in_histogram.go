package array

func largestRectangleArea(heights []int) int {
	maxArea, stack, height := 0, []int{}, 0
	for i := 0; i <= len(heights); i++ {
		if i == len(heights) {
			height = 0
		} else {
			height = heights[i]
		}
		if len(stack) == 0 || height >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			length := 0
			if len(stack) == 0 {
				length = i
			} else {
				length = i - 1 - stack[len(stack)-1]
			}
			maxArea = max84(maxArea, heights[tmp]*length)
			i--
		}
	}
	return maxArea
}

func max84(x, y int) int {
	if x > y {
		return x
	}
	return y
}
