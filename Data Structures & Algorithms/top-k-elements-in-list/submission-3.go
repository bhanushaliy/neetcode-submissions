func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for i := 0; i < len(nums); i++{
		count[nums[i]]++
	}

	for num, cnt := range count{
		freq[cnt] = append(freq[cnt], num)
	}

	res := []int{}
	for i := len(freq)-1; i > 0; i--{
		for _, num := range freq[i]{
			res = append(res, num)
			if len(res) == k{
				return res
			}
		}
	}
	return res
}
