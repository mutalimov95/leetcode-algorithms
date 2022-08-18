package longestPalindromicSubstring

func longestPalindrome(s string) string {
	lenString := len(s)
	switch lenString {
	case 0:
		return ""
	case 1:
		return s
	}
	largest := string(s[0])
	for i := 0; i < lenString; i++ {
		odd := findAround(s, i, i)
		even := findAround(s, i, i+1)

		if len(odd) > len(largest) {
			largest = odd
		}
		if len(even) > len(largest) {
			largest = even
		}
	}

	return largest
}

func findAround(s string, left, right int) string {
	var cur string
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			cur = s[left : right+1]
			left -= 1
			right += 1
		} else {
			break
		}
	}
	return cur
}
