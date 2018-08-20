/*
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

package main

import (
	"fmt"
)

func main() {
	// 1
	result := generate(5)
	fmt.Println(result)
}

func generate(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)

		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				result[i][j] = 1
			} else {
				result[i][j] = result[i-1][j-1] + result[i-1][j]
			}
		}
	}

	return result
}
