package groupAnagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)

	store := make(map[string][]string)

	if len(strs) == 1 {
		return append(res, strs)
	}

	for _, s := range strs {
		runeKey := []rune(s)
		sort.Slice(runeKey, func(i, j int) bool {
			return runeKey[i] < runeKey[j]
		})
		key := string(runeKey)
		if v, ok := store[key]; ok {
			store[key] = append(v, s)
		} else {
			store[key] = []string{s}
		}
	}

	for _, v := range store {
		res = append(res, v)
	}

	return res
}
