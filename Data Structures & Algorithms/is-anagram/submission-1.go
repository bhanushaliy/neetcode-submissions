func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	sArr := [26]int{}
	tArr := [26]int{}

	for i := 0; i < len(s); i++{
		sArr[s[i]-'a']++
		tArr[t[i]-'a']++
	}
	return sArr == tArr
}
