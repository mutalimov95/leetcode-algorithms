package threeSum

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	//output := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	fmt.Println(threeSum(input))
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	store := make(map[int]int)
	keys := make([]int, 0, len(nums))
	for _, v := range nums {
		if _, ok := store[v]; ok {
			store[v] += 1
		} else {
			store[v] = 1
		}
	}
	for k := range store {
		keys = append(keys, k)
	}

	duplicateStore := make(map[[3]int]bool)
	for i := 0; i < len(keys); i++ {
		for j := i; j < len(keys); j++ {
			k1 := keys[i]
			v1 := store[k1]

			k2 := keys[j]
			v2 := store[k2]

			k3 := 0 - k1 - k2

			if _, ok := store[k3]; !ok {
				continue
			}

			if k1+k2+k3 != 0 {
				continue
			}

			if k1 == k2 && k2 == k3 && v1 < 3 {
				continue
			} else if k1 == k2 && v1 < 2 {
				continue
			} else if k2 == k3 && v2 < 2 {
				continue
			} else if k1 == k3 && v1 < 2 {
				continue
			}

			tirplet := []int{k1, k2, k3}
			sort.Ints(tirplet)
			arrayTriplet := [3]int{tirplet[0], tirplet[1], tirplet[2]}
			if duplicateStore[arrayTriplet] {
				continue
			}
			duplicateStore[arrayTriplet] = true

			res = append(res, []int{k1, k2, k3})
		}
	}
	return res
}
