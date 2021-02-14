package main

import (
	"fmt"
	"strconv"
)

/*
input: integer
output: integer
rules:
	- We have a number sequence [1,10,9,12,3,4]
	- beginning with the rightmost digit of the input integer, multiply each digit by the corresponding number in the number sequence.
	- We first use the number at index 0 of the number sequence, and then increase the index of that number for each digit that we are multiplying it by from the input integer.
	- if our input integer is more than six digits long, we need to go back to the first number from numSequence.
	- in other words, we want to multiply each digit (assuming we've reversed our input number) by its corresponding index in numSequence mod 6

algorithm:
	- enter infinite loop
	- convert n to string, reverse it, assign to var reversedNum
	- initialize var newNum int
	- iterate over digits in reversedNum: for each digit,
		- find its corresponding number in num sequence and multiply
			- initialize var numToMultiply, set to the numSequence at index[index of reversedNum % 6]
		- add result to newNum
	- if newNum == n, return newNum
	- otherwise, set n to equal newNum

*/

func reverse(str string) string {
	var reversed string
	for _, val := range str {
		reversed = string(val) + reversed
	}
	return reversed
}

func thirt(n int) int {
	numSequence := []int{1, 10, 9, 12, 3, 4}
	for {
		reversedNum := reverse(strconv.Itoa(n))
		var newNum int
		for idx, numRune := range reversedNum {
			numToMultiply := numSequence[idx%6]
			num, _ := strconv.Atoi(string(numRune))
			newNum += num * numToMultiply
		}
		if newNum == n {
			return newNum
		}
		n = newNum
	}
	return 0
}

func main() {
	fmt.Println(thirt(321))
}
