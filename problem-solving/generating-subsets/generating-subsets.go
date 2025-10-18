// https://leetcode.com/problems/subsets/description/

package main

import "fmt"

func subsets(nums []int) [][]int {
	subsets := [][]int{{}}

	for _, num := range nums {
		for _, set := range subsets {

			subsets = append(subsets, append([]int{}, append(set, num)...))

		}

	}

	return subsets

}

func main() {

	fmt.Printf("%v\n", subsets([]int{1, 2, 3}))

}
