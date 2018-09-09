/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
*/

package main

import (
	"fmt"
)

func main() {
	var (
		haystack string
		needle   string
		result   int
	)

	// 1
	haystack = "hello"
	needle = "ll"
	result = strStr(haystack, needle)
	// 2
	fmt.Println(result)

	// 2
	haystack = "aaaaa"
	needle = "bba"
	result = strStr(haystack, needle)
	// -1
	fmt.Println(result)

	// 3
	haystack = ""
	needle = ""
	result = strStr(haystack, needle)
	// 0
	fmt.Println(result)

	// 4
	haystack = "mississippi"
	needle = "pi"
	result = strStr(haystack, needle)
	// 9
	fmt.Println(result)
}

func strStr(haystack string, needle string) int {
	result := -1
	needleLen := len(needle)
	haystackLen := len(haystack)

	if needle == "" || haystack == needle {
		return 0
	}

	if needleLen > haystackLen {
		return -1
	}

	for i := 0; i < haystackLen-needleLen+1; i++ {
		lastIndex := i + needleLen
		if haystack[i:lastIndex] == needle {
			result = i
			break
		}
	}

	return result
}
