func rotate(nums []int, k int) {
	n := len(nums)
	tmp := make([]int, n)

	for i := 0; i < n; i++{
		tmp[(i+k)%n] = nums[i]
	}

	for i := 0; i < n; i++{
		nums[i] = tmp[i]
	}
}
