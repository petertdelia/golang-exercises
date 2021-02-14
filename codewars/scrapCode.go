package main

import "fmt"

func changeSlice(slice *[]int) {
	*slice = append(*slice, 4)
}

func main() {
	slice := []int{2, 3, 4}
	fmt.Println(slice)
	changeSlice(&slice)
	fmt.Println(slice)

}
