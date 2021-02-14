// Take an integer n (n >= 0) and a digit d (0 <= d <= 9) as an integer. Square all numbers k (0 <= k <= n) between 0 and n. Count the numbers of digits d used in the writing of all the k**2. Call nb_dig (or nbDig or ...) the function taking n and d as parameters and returning this count.
// Examples:

// n = 10, d = 1, the k*k are 0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100
// We are using the digit 1 in 1, 16, 81, 100. The total count is then 4.

// nb_dig(25, 1):
// the numbers of interest are
// 1, 4, 9, 10, 11, 12, 13, 14, 19, 21 which squared are 1, 16, 81, 100, 121, 144, 169, 196, 361, 441
// so there are 11 digits `1` for the squares of numbers between 0 and 25.

package main

import (
	"fmt"
	"strconv"
)

func countDigit(n int, d int) int {
	/*
		- initialize var count to 0
		- iterate from 0 to n. count the number of digits d in the resulting square of k.
			- convert number to string.
			- iterate over string and see if any of them match d. If they do, increment count
		- return count
	*/
	count := 0
	for k := 0; k <= n; k++ {
		square := strconv.Itoa(k * k)
		for _, val := range square {
			if string(val) == strconv.Itoa(d) {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(countDigit(25, 1))
}
