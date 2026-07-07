func majorityElement(nums []int) int {
    n := len(nums)
	majority := n/2
	freq := make(map[int]int)

	if n < 1{
		return 0
	}

	for i := 0; i < n; i++{
		if val, ok := freq[nums[i]]; ok{
			if val >= majority{
				return nums[i]
			}
		}
		freq[nums[i]]++
	}
	return nums[0]
}
