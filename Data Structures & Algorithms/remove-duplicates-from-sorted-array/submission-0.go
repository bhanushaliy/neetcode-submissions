func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++{
		if nums[i] == nums[j]{
			continue
		}else{
			nums[i+1] = nums[j]
			i++
		}
	}
	return i+1
}
