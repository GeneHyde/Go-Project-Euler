//Problem 14 is to find the longest Collatz Sequence under 1 million
//a Collatz sequence being taking an even number and dividing it by 2 and
//an odd number and multiplying 3 and adding 1 until the number reaches 1

package main

import "fmt"

func Collatz(n, len int) int {
	if n <= 1 {
		return len
	} else if n % 2 == 0{
		return Collatz(n/2, len+1)
	} else {
		return Collatz((3 * n) + 1, len+1)
	}
}

func main() {
	longest := []int{0,0}
	temp := 0
	for i := 0; i < 1e6 ; i++ {
		temp = Collatz(i, 0)
		if temp > longest[1] {
			longest[0] = i
			longest[1] = temp
		}
	}
	fmt.Println(longest)
}