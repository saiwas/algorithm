/*
Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.

Example:

Input: s = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: the subarray [4,3] has the minimal length under the problem constraint.
Follow up:
If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n).
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
	inputData = []int{2, 3, 1, 2, 4, 3}
	result = minSubArrayLen(7, inputData)
	// 2
	fmt.Println(result)

	// 2
	inputData = []int{}
	result = minSubArrayLen(7, inputData)
	// 0
	fmt.Println(result)

	// 3
	inputData = []int{1, 4, 4}
	result = minSubArrayLen(4, inputData)
	// 1
	fmt.Println(result)

	// 4
	inputData = []int{7}
	result = minSubArrayLen(7, inputData)
	// 1
	fmt.Println(result)
}

func minSubArrayLen(s int, nums []int) int {
	numsLen := len(nums)
	length, j := 0, 0

	if numsLen < 1 {
		return 0
	}

	sum := nums[j]
	i := j + 1

	for j < i {
		if sum < s {
			if i == numsLen {
				break
			}
			sum += nums[i]
			i++
		} else {
			tmp := i - j
			if length > tmp || length == 0 {
				length = tmp
			}
			sum -= nums[j]
			j++
		}
	}

	return length
}
