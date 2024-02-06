package main

import (
	"fmt"	
)

func main(){

	//forma 1
	var a [3][3]int
	a[0][0] = 1
	a[0][1] = 2
	a[0][2] = 3
	a[1][0] = 4
	a[1][1] = 5
	a[1][2] = 6
	a[2][0] = 7
	a[2][1] = 8
	a[2][2] = 9
	fmt.Println(a)

	//forma 2
	var b = [3][2]int{
		{1,2},
		{3,4},
		{5,6},

	}
	fmt.Println(b)
	fmt.Println(b[2][1])

	
	

}
