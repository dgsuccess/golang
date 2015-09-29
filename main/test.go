package main

import (
	"fmt"
)

func main() {
	a := [3][2][2]int{
		1: [2][2]int{
			[2]int{1, 2},
			[2]int{3, 4},
		},
		0: [2][2]int{
			[2]int{1, 2},
			[2]int{3, 4},
		},
		2: [2][2]int{
			[2]int{1, 2},
			[2]int{3, 4},
		},
	}
	fmt.Println(a)

}
