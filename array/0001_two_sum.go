package array

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		another := target - num
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[num] = i
	}
	return nil
}
