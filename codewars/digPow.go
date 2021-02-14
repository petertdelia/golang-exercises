package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
Some numbers have funny properties. For example:

    89 --> 8¹ + 9² = 89 * 1

    695 --> 6² + 9³ + 5⁴= 1390 = 695 * 2

    46288 --> 4³ + 6⁴+ 2⁵ + 8⁶ + 8⁷ = 2360688 = 46288 * 51

Given a positive integer n written as abcd... (a, b, c, d... being digits) and a positive integer p

    we want to find a positive integer k, if it exists, such as the sum of the digits of n taken to the successive powers of p is equal to k * n.

In other words:

    Is there an integer k such as : (a ^ p + b ^ (p+1) + c ^(p+2) + d ^ (p+3) + ...) = n * k

If it is the case we will return k, if not return -1.

Note: n and p will always be given as strictly positive integers.

digPow(89, 1) should return 1 since 8¹ + 9² = 89 = 89 * 1
digPow(92, 1) should return -1 since there is no k such as 9¹ + 2² equals 92 * k
digPow(695, 2) should return 2 since 6² + 9³ + 5⁴= 1390 = 695 * 2
digPow(46288, 3) should return 51 since 4³ + 6⁴+ 2⁵ + 8⁶ + 8⁷ = 2360688 = 46288 * 51

input: two ints
output: int

algo:
	- initialize var sum int
	- split n into its digits
		- convert n to string
		- split n into array of string
		- iterate over array of strings
			- for each rune, convert to string, convert to int
			- multiply digit by p and add to sum
			- increment p
	- check if sum % n == 0
		- if true, return sum / n
*/

func digPow(n, p int) int {
	var sum float64
	numStr := strings.Split(strconv.Itoa(n), "")
	for _, val := range numStr {
		num, _ := strconv.Atoi(string(val))
		sum += math.Pow(float64(num), float64(p))
		p++
	}
	if int(sum)%n == 0 {
		return int(sum) / n
	}
	return -1
}

func main() {
	fmt.Println(digPow(46288, 3))
}
