func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freqArr := make([][]int, len(nums)+1)
	res := []int{}

	for i := 0; i < len(nums); i++{
		count[nums[i]]++
	}

	for num, cnt := range count{
		freqArr[cnt] = append(freqArr[cnt], num)
	}

	for i := len(freqArr)-1; i > 0; i--{
		for _, num := range freqArr[i]{
			res = append(res, num)
			if len(res) == k{
				return res
			}
		}
	}
	return res
}