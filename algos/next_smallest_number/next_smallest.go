package main

import (
	"fmt"
	"sort"
)

func nextSmallestNumber(nums1 []int, target int) int {

	nums := nums1
	sort.Ints(nums)

	for _, item := range nums {
		if target <= item {
			return item
		}
	}

	return -1

}
func main() {

	favCards := []int{3, 4, 6, 9, 10, 12, 14, 15, 17, 19, 21}
	target := 12

	solution := nextSmallestNumber(favCards, target)

	fmt.Printf("Next smallest number %d .\n", solution)
}
