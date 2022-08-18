package longestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {
	var counter, maxLen int
	runeS := []rune(s)
	cache := make(map[rune]bool)
	for i, c1 := range runeS {
		cache[c1] = true
		counter = 1
		for j := i + 1; j < len(s); j++ {
			c2 := runeS[j]
			if cache[c2] {
				clear(cache)
				break
			}
			cache[c2] = true
			counter += 1
		}
		if counter > maxLen {
			maxLen = counter
		}
	}
	return maxLen
}

func clear(m map[int32]bool) {
	for k := range m {
		delete(m, k)
	}
}
