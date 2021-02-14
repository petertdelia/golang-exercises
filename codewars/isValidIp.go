package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
input: string of digits and periods
output: boolean
requirements:
	- that the input value is a string
	- that it consists of four period-separated digit strings of max length 3
	- that each digit string does NOT begin with '0'
	- that each digit string is 0 <= digitString <= 255

algorithm:
	- test that the string contains only d+\.d+\.d+\.d+ -- if it doesn't, return false
	- split the string on '.', assign result to var splitStr
	- iterate over splitStr
		- for each value, check if the first char is '0' using strings.HasPrefix. If it is AND the length of the string split into all chars is greater than 1, return false
		- then convert the value to a number. Check if it is <= 255 and >= 0. If it is not, return false.
	- after iterating, return true
*/

func isValidIP(ip string) bool {
	re := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)
	if !re.Match([]byte(ip)) {
		return false
	}
	splitStr := strings.Split(ip, ".")
	for _, val := range splitStr {
		if strings.HasPrefix(val, "0") && len(strings.Split(val, "")) > 1 {
			return false
		}
		num, _ := strconv.Atoi(val)
		if num > 255 || num < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isValidIP("03.1.2.4"))
	fmt.Println(1/2 + 2/3 + 3/4 + 4/5)
}
