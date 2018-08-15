/*
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.
*/

package main

import (
	"fmt"
)

func main() {
	var (
		inputData []int
	)

	// 1
	inputData = []int{0, 1, 0, 3, 12}
	moveZeroes(inputData)
	// [1,3,12,0,0]
	fmt.Println(inputData)

	// 2
	inputData = []int{1, 2, 0, 0, 3, 12}
	moveZeroes(inputData)
	// [1,2,3,12,0,0]
	fmt.Println(inputData)

	// 3
	inputData = []int{2, 1}
	moveZeroes(inputData)
	// [2,1]
	fmt.Println(inputData)
}

func moveZeroes(nums []int) {
	sentinel := 0

	for index, value := range nums {
		if value != 0 {
			if sentinel != index {
				nums[sentinel], nums[index] = nums[index], nums[sentinel]
			}
			sentinel++
		}
	}
}
