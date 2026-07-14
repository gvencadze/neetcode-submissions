func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]int)

    for i := 0; i < len(nums); i++ {
		m[nums[i]]++

		if m[nums[i]] > (len(nums) / 2) {
			return nums[i]
		}
	}

	return 0
}
