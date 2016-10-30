//Problem 6 is to find the difference between the square of the sums of the numbers
//between 1 and 100 and the sum of the squares between 1 and 100

package main

import (
	"fmt"
	"math"
)

func SumSquare(n int) int{
	ret := 0
	for i:= 1; i <= n; i++ {
		ret = ret + int(math.Pow(float64(i),2))
	}
	return ret
}

func SquareSum(n int) int{
	ret := 0
	for i:= 1; i <= n; i++ {
		ret = ret + i
	}
	return int(math.Pow(float64(ret), 2))
}

func main() {
	fmt.Println(SquareSum(100) - SumSquare(100))
}