package main

import (
	"fmt"
	"strings"
)

// In this Kata, we are going to reverse a string while maintaining the spaces (if any) in their original place.

// For example:

// solve("our code") = "edo cruo"
// -- Normal reversal without spaces is "edocruo".
// -- However, there is a space at index 3, so the string becomes "edo cruo"

// solve("your code rocks") = "skco redo cruoy".
// solve("codewars") = "srawedoc"

// input: string
// output: string
// rules:
// 	- preserve spaces in their original indexes

// algo:
// 	- initialize var reversed []string
// 	- iterate through string
// 		-
func reverseWithSpaces(str string) string {
	var reversed []string
	var spaces []int
	for idx, val := range str {
		if val != ' ' {
			reversed = append([]string{string(val)}, reversed...)
		} else {
			spaces = append(spaces, idx)
		}
	}
	for _, val := range spaces {
		reversed = append(reversed, " ")
		copy(reversed[val+1:], reversed[val:])
		reversed[val] = " "
	}
	return strings.Join(reversed, "")
}

func main() {
	fmt.Println(reverseWithSpaces("he ll o"))
}
