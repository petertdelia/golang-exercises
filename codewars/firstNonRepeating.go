package main

import (
	"fmt"
	"strings"
)

/*
Write a function named first_non_repeating_letter that takes a string input, and returns the first character that is not repeated anywhere in the string.

For example, if given the input 'stress', the function should return 't', since the letter t only occurs once in the string, and occurs first in the string.

As an added challenge, upper- and lowercase letters are considered the same character, but the function should return the correct case for the initial letter. For example, the input 'sTreSS' should return 'T'.

If a string contains all repeating characters, it should return an empty string ("") or None -- see sample tests.

input: string
output: string, that is the first character that is not repeated in the string.
	- counting a character as repeated ignores case;
	- however, the returned character should be in the correct case -- that is, the case that it appeared in the original string.
	- if there is no unrepeated char, return empty string.

What do we need to keep track of?
	- the index of each char
	- after we go through once, counting each char, we need to find the char that was only counted once with the lowest index num.

data structures:
	- go map/hash table structure to count chars

algorithm:
	- initialize maps, charCount, filteredCharCount
	- iterate through lowercased string
		- for each char, increment charCount at that char
	- after iterating, filter charCount, , looking only for those chars that have count of 1
		-
	- iterate through (lowercase) string again, checking each char to see if it exists in filteredCharCount.
		- If it does, get the char's index and return the original string at that index.
		- if it does not, after iterating return empty string.

Q: how do we get an array of runes? Strings convert to array of runes natively.
*/

func firstNonRepeating(str string) string {
	charCount := make(map[string]int)
	lowerCaseRunes := []rune(strings.ToLower(str))
	for _, val := range lowerCaseRunes {
		charCount[string(val)]++
	}
	for idx, val := range lowerCaseRunes {
		if charCount[string(val)] == 1 {
			return string([]rune(str)[idx])
		}
	}

	return ""
}

func main() {
	fmt.Println(firstNonRepeating("sTreSS"))
}
