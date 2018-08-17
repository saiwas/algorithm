/*
Given a matrix of M x N elements (M rows, N columns), return all elements of the matrix in diagonal order as shown in the below image.

Example:
Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output:  [1,2,4,7,5,3,6,8,9]

Note:
The total number of elements of the given matrix will not exceed 10,000.
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
	result = findDiagonalOrder(inputData)
	// []
	fmt.Println(result)

	// 2
	inputData = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	result = findDiagonalOrder(inputData)
	// [1,2,4,7,5,3,6,8,9]
	fmt.Println(result)

	// 3
	inputData = [][]int{
		{6, 9, 7},
	}
	result = findDiagonalOrder(inputData)
	// []
	fmt.Println(result)

	// 4
	inputData = [][]int{
		{6}, {9}, {7},
	}
	result = findDiagonalOrder(inputData)
	// []
	fmt.Println(result)

	// 5
	inputData = [][]int{
		{1},
	}
	result = findDiagonalOrder(inputData)
	// [1]
	fmt.Println(result)

	// 6
	inputData = [][]int{
		{3}, {2},
	}
	result = findDiagonalOrder(inputData)
	// [3,2]
	fmt.Println(result)
}

func findDiagonalOrder(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0 {
		return nil
	}
	columns := len(matrix[0])
	if columns == 0 {
		return nil
	}
	result := make([]int, 0, rows*columns)
	for rowIndex, columnIndex, d := 0, 0, 0; rowIndex+columnIndex <= (rows-1)+(columns-1); {
		result = append(result, matrix[rowIndex][columnIndex])
		if d == 0 {
			if rowIndex == 0 {
				if columnIndex == columns-1 {
					rowIndex, d = rowIndex+1, 1
				} else {
					columnIndex, d = columnIndex+1, 1
				}
			} else if columnIndex == columns-1 {
				rowIndex, d = rowIndex+1, 1
			} else {
				rowIndex, columnIndex = rowIndex-1, columnIndex+1
			}
		} else {
			if rowIndex == rows-1 {
				columnIndex, d = columnIndex+1, 0
			} else if columnIndex == 0 {
				rowIndex, d = rowIndex+1, 0
			} else {
				rowIndex, columnIndex = rowIndex+1, columnIndex-1
			}
		}
	}
	return result
}
