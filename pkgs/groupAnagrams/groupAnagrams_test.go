package groupAnagrams

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	wants := [][][]string{
		{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		{{""}},
		{{"a"}},
	}
	inputs := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
		{""},
		{"a"},
	}
	for i, input := range inputs {
		want := wants[i]
		res := groupAnagrams(input)
		if !reflect.DeepEqual(res, want) {
			t.Errorf("On string \"%v\" want %v, but got %v", input, want, res)
		}
	}
}
