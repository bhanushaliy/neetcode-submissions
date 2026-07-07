func sortColors(nums []int) {
    l := 0
	r := len(nums)-1
	i := 0

	for i <= r{
		if nums[i] == 2{
			nums[i], nums[r] = nums[r], nums[i]
			r--
		}else if nums[i] == 0{
			nums[i], nums[l] = nums[l], nums[i]
			i++
			l++
		}else{
			i++
		}
	}
}
