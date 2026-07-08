func longestConsecutive(nums []int) int {
	seen := make(map[int]bool)

	for _, num := range nums{
		seen[num] = true
	}

	longest := 0

	for _, ele := range nums{
		if seen[ele-1]{
			continue
		}

		current := ele
		length := 1


		for seen[current+1]{
			current++
			length++
		}

		if length > longest{
			longest = length
		}
	}
	return longest
}