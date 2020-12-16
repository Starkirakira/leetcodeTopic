package lc

import "strings"

func WordPattern(pattern string, s string) bool {
	wordList := map[string]byte{}
	sList := map[byte]string{}
	word :=strings.Split(s, " ")

	if len(pattern) != len(word) {
		return false
	}

	for i,j := range word {
		ch := pattern[i]

		if wordList[j] > 0 && wordList[j] != ch || sList[ch] != "" && sList[ch] != j {
			return false
		}
		wordList[j] = ch
		sList[ch] = j
	}
	return true


}
