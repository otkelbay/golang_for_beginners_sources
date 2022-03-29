package main

import "fmt"

func testArrays() {
	var matrix [][]int

	matrix = [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}

	matrix = append(matrix, []int{4, 5, 6})

	fmt.Println(matrix)
}
