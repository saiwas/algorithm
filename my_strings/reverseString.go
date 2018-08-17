/*
Write a function that takes a string as input and returns the string reversed.

Example 1:

Input: "hello"
Output: "olleh"
Example 2:

Input: "A man, a plan, a canal: Panama"
Output: "amanaP :lanac a ,nalp a ,nam A"
*/

package main

import (
	"fmt"
)

func main() {
	var (
		inputData string
		result    string
	)

	// 1
	inputData = "hello"
	result = reverseString(inputData)
	// olleh
	fmt.Println(result)

	// 2
	inputData = "A man, a plan, a canal: Panama"
	result = reverseString(inputData)
	// amanaP :lanac a ,nalp a ,nam A
	fmt.Println(result)
}

func reverseString(s string) string {
	bytes := []byte(s)
	i := 0
	j := len(bytes) - 1

	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}

	return string(bytes)
}
