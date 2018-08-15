/*
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
*/

package main

import (
	"fmt"
)

func main() {
	var (
		inputData []int
		result    []int
	)

	// 1
	inputData = []int{1, 2, 3}
	result = plusOne(inputData)
	// [1,2,4]
	fmt.Println(result)

	// 2
	inputData = []int{4, 3, 2, 1}
	result = plusOne(inputData)
	// [4,3,2,2]
	fmt.Println(result)

	// 3
	inputData = []int{9}
	result = plusOne(inputData)
	// [1,0]
	fmt.Println(result)

	// 4
	inputData = []int{1, 2, 9}
	result = plusOne(inputData)
	// [1, 3, 0]
	fmt.Println(result)

	// 5
	inputData = []int{9, 9, 9}
	result = plusOne(inputData)
	// [1, 0, 0, 0]
	fmt.Println(result)
}

func plusOne(digits []int) []int {
	lastIndex := len(digits) - 1

	digits[lastIndex]++

	for i := lastIndex; i >= 0; i-- {
		if digits[i] > 9 {
			digits[i] = 0
			if i < 1 {
				digits = append([]int{1}, digits...)
			} else {
				digits[i-1]++
			}
		} else {
			break
		}
	}

	return digits
}
