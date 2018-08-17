/*
Given a binary array, find the maximum number of consecutive 1s in this array.

Example 1:
Input: [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
    The maximum number of consecutive 1s is 3.
Note:

The input array will only contain 0 and 1.
The length of input array is a positive integer and will not exceed 10,000
*/

package main

import (
	"fmt"
)

func main() {
	var (
		inputData []int
		result    int
	)

	// 1
	inputData = []int{1, 1, 0, 1, 1, 1}
	result = findMaxConsecutiveOnes(inputData)
	// 3
	fmt.Println(result)
}

func findMaxConsecutiveOnes(nums []int) int {
	counts := 0
	maxConsecutive := 0

	for _, value := range nums {
		if value == 1 {
			counts++
			if maxConsecutive < counts {
				maxConsecutive = counts
			}
		} else {
			counts = 0
		}
	}

	return maxConsecutive
}
