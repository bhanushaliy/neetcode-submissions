func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	freq := make(map[byte]int)

	for i := 0; i < len(s); i++{
		freq[s[i]]++
		freq[t[i]]--
	}

	for _, val := range freq{
		if val != 0{
			return false
		}
	}
	return true
}
