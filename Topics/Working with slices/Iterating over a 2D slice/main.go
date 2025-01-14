package main

import "fmt"

func main() {
	// DO NOT delete the code block below!
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	var num int
	fmt.Scanln(&num)

	// Use a nested for...range loop to print each element
	// of the matrix in a new line, multiplied by 'num'

	for a, _ := range matrix {
		for b, _ := range matrix {
			fmt.Println(matrix[a][b] * num)
		}
	}
}
