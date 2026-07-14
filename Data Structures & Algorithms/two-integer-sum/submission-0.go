func twoSum(nums []int, target int) []int {
	indexes := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		for j := i+1 ; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				indexes = []int{i, j}
				break
			}
		}
	}

	return indexes
}
