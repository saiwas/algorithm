/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.
*/

package main

import (
	"fmt"
)

func main() {
	var (
		input  []string
		output string
	)

	// 1
	input = []string{"flower", "flow", "flight"}
	output = longestCommonPrefix(input)
	// "fl"
	fmt.Println(output)

	// 2
	input = []string{"dog", "racecar", "car"}
	output = longestCommonPrefix(input)
	// ""
	fmt.Println(output)

	// 3
	input = []string{}
	output = longestCommonPrefix(input)
	// ""
	fmt.Println(output)

	// 4
	input = []string{"a", "ac"}
	output = longestCommonPrefix(input)
	// "a"
	fmt.Println(output)

	// 5
	input = []string{"aa", "a"}
	output = longestCommonPrefix(input)
	// "a"
	fmt.Println(output)

	// 6
	input = []string{"ca", "a"}
	output = longestCommonPrefix(input)
	// ""
	fmt.Println(output)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		prefixLen := len(prefix)
		strLen := len(strs[i])

		if prefixLen == 0 {
			break
		}
		for j := 0; j < strLen; j++ {
			if j >= prefixLen {
				break
			}
			if strs[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}

		if len(prefix) > strLen {
			prefix = strs[i]
		}
	}

	return prefix
}
