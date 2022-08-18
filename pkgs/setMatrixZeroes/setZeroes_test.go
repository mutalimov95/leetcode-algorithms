package setMatrixZeroes

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	matrix1 := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	matrix2 := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	want1 := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}
	want2 := [][]int{
		{0, 0, 0, 0},
		{0, 4, 5, 0},
		{0, 3, 1, 0},
	}
	inputs := [][][]int{matrix1, matrix2}
	wants := [][][]int{want1, want2}
	for i, input := range inputs {
		want := wants[i]
		setZeroes(input)
		res := input
		if !reflect.DeepEqual(res, want) {
			t.Errorf("On string \"%v\" want %v, but got %v", input, want, res)
		}
	}
}
