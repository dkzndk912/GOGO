package main

import "fmt"

func main() {
	var s []int
	s = make([]int, 5)

	for _, v := range s {
		fmt.Println(v)
	}
}
