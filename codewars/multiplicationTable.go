package main

import (
	"fmt"
)

/*
algo:
	- initialize var result to a multidimensional slice
	- run for loop, using i, from 1 to size
		- initialize var row to a single dimensional slice
		- run for loop, using j, from 1 to size
			- in inner loop, append int i * j to row
		- in outer loop append row to result
	- return result
*/

func multiplicationTable(size int) (result [][]int) {
	for i := 1; i <= size; i++ {
		var row []int
		for j := 1; j <= size; j++ {
			row = append(row, i*j)
		}
		result = append(result, row)
	}
	return result
}

func main() {
	fmt.Println(multiplicationTable(3))
}
