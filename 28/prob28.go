//Problem 28 is to build a 1001 x 1001 size matrix where the numbers expand out from
//the middel starting at one (easier to picture if you see the actual problem)
//and to calculate the sum of the diagonals 

package main

import (
	"fmt"
)

func OutBound(x, y, size int) bool {
	if x >= size || x < 0 || y >= size || y < 0 {
		return false
	}
	return true
}

func MoveHorz(dist, set, x, y *int, direction int, ara [][]int) {
	for i:= 0; i < *dist && OutBound(*x, *y, len(ara)); i++ {
		ara[*y][*x] = *set
		*set += 1
		*x += direction
	}
	*dist += 1
}

func MoveVert(dist, set, x, y *int, direction int, ara [][]int) {
	for i:= 0; i < *dist && OutBound(*x, *y, len(ara)); i++ {
		ara[*y][*x] = *set
		*set += 1
		*y += direction
	}
	*dist += 1
}

func BuildArray(size int) [][]int {
	array := make([][]int, size)
	for i:=0; i<len(array); i++ {
		array[i] = make([]int, size)
	}
	horz := 1
	vert := 1
	x := size/2
	y := size/2
	set := 1
	for {
		if !OutBound(x, y, size){
			return array
		} 
		MoveHorz(&horz, &set, &x, &y, 1, array)
		MoveVert(&vert, &set, &x, &y, 1, array)
		MoveHorz(&horz, &set, &x, &y, -1, array)
		MoveVert(&vert, &set, &x, &y, -1, array)
	}
}

func Calc(ara [][]int) int {
	total := 0
	size := len(ara)
	x:= 0 
	y:= 0
	for ;x < size && y < size; {
		total += ara[y][x]
		x++ 
		y++
	}
	x = size-1 
	y = 0
	for ;x < size && y < size;{
		total += ara[y][x]
		x-- 
		y++
	}
	return total -1
}

func main() {
	ara := BuildArray(1001)
	fmt.Println(Calc(ara))
}