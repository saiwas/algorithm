/*
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		a      string
		b      string
		result string
	)

	// 1
	a = "11"
	b = "1"

	// 100
	result = addBinary(a, b)
	fmt.Println(result)

	// 2
	a = "1010"
	b = "1011"

	// 10101
	result = addBinary(a, b)
	fmt.Println(result)
}

func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	result := ""

	if aLen == 0 {
		return b
	}
	if bLen == 0 {
		return a
	}

	aInts := make([]int, aLen)
	bInts := make([]int, bLen)
	i := -1
	j := -1
	var tmp string

	for _, s := range a {
		tmp = string(s)
		i++
		aInts[i], _ = strconv.Atoi(tmp)
	}

	for _, s := range b {
		tmp = string(s)
		j++
		bInts[j], _ = strconv.Atoi(tmp)
	}

	carries := 0
	bj := 0

	for i := i; i >= 0; i-- {
		if j < 0 {
			bj = 0
		} else {
			bj = bInts[j]
		}

		sum := aInts[i] + bj + carries
		bit := sum % 2
		carries = sum / 2

		result = strconv.Itoa(bit) + result
		j--
	}

	for j := j; j >= 0; j-- {
		sum := bInts[j] + carries
		bit := sum % 2
		carries = sum / 2

		result = strconv.Itoa(bit) + result
	}

	if carries > 0 {
		result = "1" + result
	}

	return result
}
