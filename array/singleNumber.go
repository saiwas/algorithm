/*
Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,1]
Output: 1
Example 2:

Input: [4,1,2,1,2]
Output: 4
*/

package main

import "fmt"

func main() {
	var inputData []int
	var k int

	// 1
	inputData = []int{2, 2, 1}
	k = singleNumber(inputData)
	// 1
	fmt.Println(k)

	// 2
	inputData = []int{4, 1, 2, 1, 2}
	k = singleNumber(inputData)
	// 4
	fmt.Println(k)
}

func singleNumber(nums []int) int {
	counter := make(map[int]int)
	numsLen := len(nums)

	for i := 0; i < numsLen; i++ {
		counter[nums[i]]++
	}

	for key, value := range counter {
		if value < 2 {
			return key
		}
	}

	return -1
}
