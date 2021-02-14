/*Take 2 strings s1 and s2 including only letters from ato z. Return a new sorted string, the longest possible, containing distinct letters - each taken only once - coming from s1 or s2.
Examples:

a = "xyaabbbccccdefww"
b = "xxxxyyyyabklmopq"
longest(a, b) -> "abcdefklmopqwxy"

a = "abcdefghijklmnopqrstuvwxyz"
longest(a, a) -> "abcdefghijklmnopqrstuvwxyz"
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func twoToOne(s1, s2 string) string {
	chars := make(map[string]bool)
	for _, val := range []rune(s1) {
		chars[string(val)] = true
	}
	for _, val := range []rune(s2) {
		chars[string(val)] = true
	}
	var result string
	for val := range chars {
		result += val
	}
	splitStr := strings.Split(result, "")
	sort.Strings(splitStr)
	return strings.Join(splitStr, "")
}

func main() {
	fmt.Println(twoToOne("hello", "me"))
}
