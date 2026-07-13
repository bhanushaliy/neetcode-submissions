func groupAnagrams(strs []string) [][]string {
	n := len(strs)
	resArr := [][]string{}
	seen := make([]bool, n)

	for i := 0; i < n; i++{
		if seen[i]{
			continue
		}
		group := []string{strs[i]}
		seen[i] = true
		for j := i+1; j < n; j++{
			if !seen[j] && isAnagram(strs[i], strs[j]){
				group = append(group, strs[j])
				seen[j] = true
			}
		}
		resArr = append(resArr, group)
	}
	return resArr
}

func isAnagram(str1 string, str2 string) bool{
	if len(str1) != len(str2){
		return false
	}

	countStr1 := [26]int{}
	countStr2 := [26]int{}

	for i := 0; i < len(str1); i++{
		countStr1[str1[i]-'a']++
		countStr2[str2[i]-'a']++
	}

	return countStr1 == countStr2
}