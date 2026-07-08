func mergeAlternately(word1 string, word2 string) string {
	minWord := ""
	maxWord := ""
	if len(word1) < len(word2){
		minWord = word1
		maxWord = word2
	}else{
		minWord = word2
		maxWord = word1
	}

	res := ""

	for i := 0; i < len(minWord); i++{
		res += string(word1[i])
		res += string(word2[i])
	}
	res += maxWord[len(minWord):]
	return res
}
