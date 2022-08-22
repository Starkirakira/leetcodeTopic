package lc

import "unicode"

func FindWords(words []string) (ans []string) {
	const index = "12210111011122000010020202"
loop:
	for _, word := range words {
		idx := index[unicode.ToLower(rune(word[0]))-'a']
		for _, letter := range word[1:] {
			if index[unicode.ToLower(rune(letter))-'a'] != idx {
				continue loop
			}
		}
		ans = append(ans, word)

	}
	return

}
