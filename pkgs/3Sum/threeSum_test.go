package threeSum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	wants := [][][]int{
		{
			{-1, -1, 2},
			{-1, 0, 1},
		},
		{},
		{
			{0, 0, 0},
		},
	}
	inputs := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{0, 1, 1},
		{0, 0, 0},
	}
	for i, input := range inputs {
		want := wants[i]
		res := threeSum(input)
		if !reflect.DeepEqual(res, want) {
			t.Errorf("On string \"%v\" want %v, but got %v", input, want, res)
		}
	}
}
