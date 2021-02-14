/* Everyone knows passphrases. One can choose passphrases from poems, songs, movies names and so on but frequently they can be guessed due to common cultural references. You can get your passphrases stronger by different means. One is the following:

choose a text in capital letters including or not digits and non alphabetic characters,

    shift each letter by a given number but the transformed letter must be a letter (circular shift),
    replace each digit by its complement to 9,
    keep such as non alphabetic and non digit characters,
    downcase each letter in odd position, upcase each letter in even position (the first character is in position 0),
    reverse the whole result.

Example:

your text: "BORN IN 2015!", shift 1

1 + 2 + 3 -> "CPSO JO 7984!"

4 "CpSo jO 7984!"

5 "!4897 Oj oSpC"

With longer passphrases it's better to have a small and easy program. Would you write it?
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func playPass(s string, n int) string {
	/*
		input: string, int
		output: string
		rules
			- shift each letter rightward by the input int, wrapping if necessary
			- replace digits by (9 - digit)
			- retain non alphanumeric chars
			- downcase alphabetic chars at odd indexes, upcase at even indexes.
			- reverse the result
			implicit:
				- input alphabetic chars will be uppercase
				- test cases involve wrapping alphabet

		data structures:
			- string
			- slice of strings (split), iteration

		algo:
			- declare var alphabet, string of all alphabetic chars
			- declare var digits, string of all digits
			- declare var result, string
			- iterate over input string
				- for each char, check if it exists in alphabet
					- if it does, find its index within alphabet, add n to index, get char at that index in alphabet. Check its index within input string. If it is odd, downcase it and prepend to result. Otherwise, upcase it and prepend to result.
				- check if char exists in digits
					- if it does, convert to int, subtract result from 9, convert back to string, and prepend to result.
				- otherwise, prepend to result
	*/
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	digits := "0123456789"
	var result string

	for idx, val := range strings.Split(s, "") {
		if strings.Contains(alphabet, strings.ToLower(val)) {
			index := strings.Index(alphabet, strings.ToLower(val))
			newIndex := (index + n) % 26
			char := string(alphabet[newIndex])
			if idx%2 == 1 {
				char = strings.ToLower(char)
			} else {
				char = strings.ToUpper(char)
			}
			result = char + result
		} else if strings.Contains(digits, val) {
			num, _ := strconv.Atoi(val)
			newNum := 9 - num
			result = strconv.Itoa(newNum) + result
		} else {
			result = val + result
		}
	}
	return result
}

func main() {
	fmt.Println(playPass("BORN IN 2015!", 1))
}
