func sortArray(nums []int) []int {
    for i := 0; i < len(nums)-1; i++{
		j := i+1
		for j < len(nums){
			if nums[j] < nums[i]{
				nums[j], nums[i] = nums[i], nums[j]
			}
			j++
		}
	}
	return nums
}
