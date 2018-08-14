/*
Given an array, rotate the array to the right by k steps, where k is non-negative.

Example 1:

Input: [1,2,3,4,5,6,7] and k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: [-1,-100,3,99] and k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
Note:

Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
Could you do it in-place with O(1) extra space?
*/

package main

import "fmt"

func main() {
	var inputData []int
	var k int

	// 1
	inputData = []int{1, 2, 3, 4, 5, 6, 7}
	k = 3
	rotate(inputData, k)
	// [5,6,7,1,2,3,4]
	fmt.Println(inputData)

	// 2
	inputData = []int{-1, -100, 3, 99}
	k = 2
	rotate(inputData, k)
	// [3,99,-1,-100]
	fmt.Println(inputData)

	// 3
	inputData = []int{-1}
	k = 2
	rotate(inputData, k)
	// [-1]
	fmt.Println(inputData)

	// 4
	inputData = []int{1, 2}
	k = 3
	rotate(inputData, k)
	// [2,1]
	fmt.Println(inputData)
}

func rotate(nums []int, k int) {
	n := len(nums)

	if n < 2 {
		return
	}
	m := n - k%n

	move(nums, 0, m-1)
	move(nums, m, n-1)
	move(nums, 0, n-1)
}

func move(nums []int, from int, to int) {
	var tmp int

	for from < to {
		tmp = nums[from]
		nums[from] = nums[to]
		nums[to] = tmp
		from++
		to--
	}
}
