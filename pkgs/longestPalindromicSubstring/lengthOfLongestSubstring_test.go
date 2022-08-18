package longestPalindromicSubstring

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	wants := []string{"", "a", "bab", "bb", "bb", "aaaa"}
	inputs := []string{"", "a", "babad", "cbbd", "bb", "aaaa"}
	for i, input := range inputs {
		want := wants[i]
		res := longestPalindrome(input)
		if res != want {
			t.Errorf("On string \"%v\" want %v, but got %v", input, want, res)
		}
	}
}
