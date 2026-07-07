func hasDuplicate(nums []int) bool {
    freq := make(map[int]bool)

	for i := 0; i < len(nums); i++{
		if freq[nums[i]]{
			return true
		}
		freq[nums[i]] = true
	}
	return false
}
