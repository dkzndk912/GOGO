package main

import "fmt"

func main() {
	// var s []int
	// s = make([]int, 5)

	// s := make([]int, 5)

	s := []int{0, 0, 0, 0, 0}

	s[4] = 99
	s[2] = 91

	for _, v := range s {
		fmt.Println(v)
	}
}
