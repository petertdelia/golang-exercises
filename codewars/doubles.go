package main

import (
	"fmt"
	"math"
)

/*
algo:
	- initialize k, n as float equivalents of maxk, maxn
	- initialize velocity
	- enter double for loop that iterates up to maxk, then maxn
		- in the inner loop, calculate the velocity of the current numbers:
			- velocity += 1.0 / countk * math.Pow(countn + 1, 2.0 * countk)
*/

func doubles(maxk, maxn int) float64 {
	k, n := float64(maxk), float64(maxn)
	var velocity float64
	for countk := 1.0; countk <= k; countk++ {
		for countn := 1.0; countn <= n; countn++ {
			velocity += 1.0 / (countk * math.Pow(countn+1.0, 2.0*countk))
		}
	}

	return velocity
}

func main() {
	fmt.Println(doubles(10, 100))
}
