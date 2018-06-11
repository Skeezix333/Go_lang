package main

import (
	"fmt"
)

func main() {
	mainSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, s := range mainSlice {
		if (s % 2) == 0 {
			fmt.Println(s, " is even")
		} else {
			fmt.Println(s, " is odd")
		}
	}
}
