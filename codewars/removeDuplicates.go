package main

import "fmt"

// In this Kata, you will be given an array of strings and your task is to remove all consecutive duplicate letters from each string in the array.

// For example:

//     dup(["abracadabra","allottee","assessee"]) = ["abracadabra","alote","asese"].

//     dup(["kelless","keenness"]) = ["keles","kenes"].

// Strings will be lowercase only, no spaces. See test cases for more examples.

/*
input: array of strings
output: array of strings
requirements:
	- collapse any consecutive duplicate letters from each string in the array
	- each string contains only lowercase letters
	- assuming returning a new array.
algo:
	- initialize result, empty string array
	- iterate over input array
		- initialize var newStr
		- convert string to array of runes
		- set lastVal to equal the first element in the rune array
		- iterate over input string, starting at index 1
			- check if the current value is different from last value
				- if it is different, then add lastVal to newStr and set lastVal to equal currentVal
			- otherwise, do nothing
		- after iterating over string, add lastVal to newStr
		- append newStr to result
*/

func dup(arr []string) []string {
	result := []string{}
	for _, str := range arr {
		runes := []rune(str)
		var newStr string
		lastVal := string(runes[0])
		for _, currentVal := range runes[1:] {
			currentVal := string(currentVal)
			if lastVal != currentVal {
				newStr += lastVal
				lastVal = currentVal
			}
		}
		newStr += lastVal
		result = append(result, newStr)
	}

	return result
}

func main() {
	fmt.Println(dup([]string{"abracadabra", "allottee", "assessee"}))
}
