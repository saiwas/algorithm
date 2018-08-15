/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		inputData []int
		k         int
		result    []int
	)

	// 1
	inputData = []int{2, 7, 11, 15}
	k = 9
	result = twoSum(inputData, k)
	// [0, 1]
	fmt.Println(result)

	// 2
	inputData = []int{3, 2, 4}
	k = 6
	result = twoSum(inputData, k)
	// [1,2]
	fmt.Println(result)

	// 3
	inputData = []int{3, 4, 3}
	k = 6
	result = twoSum(inputData, k)
	// [0,2]
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	var indexes []int
	begin := 0
	end := len(nums) - 1

	shawdow := make([]int, end+1)
	copy(shawdow, nums)
	sort.Ints(shawdow)

	for begin < end {
		currentSum := shawdow[begin] + shawdow[end]

		if currentSum == target {
			break
		} else if currentSum < target {
			begin++
		} else {
			end--
		}
	}

	if begin < end {
		for index, value := range nums {
			if shawdow[begin] == value || shawdow[end] == value {
				indexes = append(indexes, index)
			}
		}
	}

	return indexes
}
