package main

import "fmt"

// Write a function that when given a number >= 0, returns an Array of ascending length subarrays.

// pyramid(0) => [ ]
// pyramid(1) => [ [1] ]
// pyramid(2) => [ [1], [1, 1] ]
// pyramid(3) => [ [1], [1, 1], [1, 1, 1] ]

/*

algo:
	- initialize var result to a two-dimensional array
	- iterate count times, where 0 < count <= n
		- in each iteration, create a new array, inner;
			- iterate ones times, where 0 < ones <= count. In each iteration, append a 1 to the array.
		- append the array to result
	- return result

*/

func pyramid(n int) [][]int {
	result := [][]int{}
	for count := 1; count <= n; count++ {
		innerArr := []int{}
		for ones := 1; ones <= count; ones++ {
			innerArr = append(innerArr, 1)
		}
		result = append(result, innerArr)
	}
	return result
}

func main() {
	fmt.Println(pyramid(5))
}
