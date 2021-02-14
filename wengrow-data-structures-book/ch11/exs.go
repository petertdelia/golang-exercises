package main

import (
	"fmt"
)

func main() {
	fmt.Println(evenNums([]int{2, 4, 1, 2, 3, 4, 5, 6}))
	fmt.Println(triangleNums(7))
	fmt.Println(indexOfX("abcdexjfk", 0))
}

func indexOfX(str string, index int) int {
	if string(str[0]) == "x" {
		return index
	} else {
		return indexOfX(str[1:], index+1)
	}
}

func triangleNums(n int) int {
	if n == 1 {
		return 1
	}

	return n + triangleNums(n-1)
}

func evenNums(nums []int) []int {
	if len(nums) == 1 {
		if nums[0]%2 == 0 {
			return nums
		} else {
			return []int{}
		}
	}
	if nums[0]%2 == 0 {
		return append([]int{nums[0]}, evenNums(nums[1:])...)
	} else {
		return evenNums(nums[1:])
	}
}
