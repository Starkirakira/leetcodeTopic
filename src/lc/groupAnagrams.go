package lc

import "sort"

func GroupAnagrams(strs []string) [][]string {
	chars := make(map[string][]string)
	for _, j := range strs {
		key := []byte(j)
		sort.Slice(key, func(i, j int) bool {
			return key[i] < key[j]
		})
		chars[string(key)] = append(chars[string(key)], j)
	}

	res := [][]string{}
	for _, v := range chars {
		res = append(res, v)
	}
	return res
}
