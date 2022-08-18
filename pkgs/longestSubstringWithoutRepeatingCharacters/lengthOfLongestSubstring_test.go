package longestSubstringWithoutRepeatingCharacters

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	wants := []int{3, 1, 3}
	inputs := []string{"abcabcbb", "bbbbb", "pwwkew"}
	for i, input := range inputs {
		want := wants[i]
		res := lengthOfLongestSubstring(input)
		if res != want {
			t.Errorf("On string \"%v\" want %v, but got %v", input, want, res)
		}
	}
}
