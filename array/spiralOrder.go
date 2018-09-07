/*
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

Example 1:

Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output: [1,2,3,6,9,8,7,4,5]
Example 2:

Input:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
*/

package main

import (
	"fmt"
)

func main() {
	var (
		inputData [][]int
		result    []int
	)

	// 1
	inputData = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	result = spiralOrder(inputData)
	// [1,2,3,6,9,8,7,4,5]
	fmt.Println(result)

	// 2
	inputData = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	result = spiralOrder(inputData)
	// [1,2,3,4,8,12,11,10,9,5,6,7]
	fmt.Println(result)

	// 3
	inputData = [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	result = spiralOrder(inputData)
	// [1,2,3,4,5,10,15,20,25,24,23,22,21,16,11,6,7,8,9,14,19,18,17,12,13]
	fmt.Println(result)
}

func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	if rows < 1 {
		return []int{}
	}
	columns := len(matrix[0])
	total := rows * columns
	result := make([]int, total)
	rowIndex, columnIndex := 0, 0
	direction := 0

	rangeIndex := map[string]int{
		"rowMinIndex": 0,
		"rowMaxIndex": rows - 1,
		"colMinIndex": 0,
		"colMaxIndex": columns - 1,
	}

	for i := 0; i < total; i++ {
		result[i] = matrix[rowIndex][columnIndex]

		// Here can change to use Finite State Machine
		switch {
		case rangeIndex["rowMinIndex"] == rowIndex && rangeIndex["colMinIndex"] == columnIndex && direction == 3:
			direction = 0
			if rangeIndex["rowMinIndex"] > 0 {
				rangeIndex["colMinIndex"]++
			}
		case rangeIndex["rowMinIndex"] == rowIndex && rangeIndex["colMaxIndex"] == columnIndex && direction == 0:
			direction = 1
			rangeIndex["rowMinIndex"]++
		case rangeIndex["rowMaxIndex"] == rowIndex && rangeIndex["colMaxIndex"] == columnIndex && direction == 1:
			direction = 2
			rangeIndex["colMaxIndex"]--
		case rangeIndex["rowMaxIndex"] == rowIndex && rangeIndex["colMinIndex"] == columnIndex && direction == 2:
			direction = 3
			rangeIndex["rowMaxIndex"]--
		}

		switch direction {
		case 0:
			columnIndex++
		case 1:
			rowIndex++
		case 2:
			columnIndex--
		case 3:
			rowIndex--
		}
	}

	return result
}
