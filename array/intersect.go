/*

Given two arrays, write a function to compute their intersection.

Example:
Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2, 2].

Note:
Each element in the result should appear as many times as it shows in both arrays.
The result can be in any order.
Follow up:
What if the given array is already sorted? How would you optimize your algorithm?
What if nums1's size is small compared to nums2's size? Which algorithm is better?
What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
*/

package main

import (
	"fmt"
)

func main() {
	var (
		nums1, nums2, result []int
	)

	// 1
	nums1 = []int{1, 2, 2, 1}
	nums2 = []int{2, 2}
	result = intersect(nums1, nums2)
	// [2,2]
	fmt.Println(result)

	// 2
	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	result = intersect(nums1, nums2)
	// [4,9]
	fmt.Println(result)

	// 3
	nums1 = []int{1, 2, 2, 1}
	nums2 = []int{1, 1}
	result = intersect(nums1, nums2)
	// [1,1]
	fmt.Println(result)
}

func intersect(nums1 []int, nums2 []int) []int {
	counter := make(map[int]bool)
	res := []int{}

	for _, value2 := range nums2 {
		counter[value2] = true
	}

	for _, value1 := range nums1 {
		if counter[value1] {
			res = append(res, value1)
		}
	}

	return res
}
