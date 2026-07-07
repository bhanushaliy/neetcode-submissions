func longestConsecutive(nums []int) int {
	if len(nums) < 1{
		return 0
	}
	maxEle := nums[0]
	minEle := math.MaxInt
	freq := make(map[int]bool)

	for _, num := range nums{
		if num > maxEle{
			maxEle = num
		}
		if num < minEle{
			minEle = num
		}
		freq[num] = true
	}

	res := 0
	count := 0
	for i := minEle; i <= maxEle; i++{
		if freq[i]{
			count++
		}else{
			if count > res{
				res = count
			}
			count = 0
		}
	}
	if count > res{
		res = count
	}
	return res
}
