package main

import "fmt"

func main() {
	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5

	primes := [3]int{2, 3, 6}
	fmt.Println(primes, primes[1])

	test := [5]bool{true, true, true}
	fmt.Println(test[3]) // boolean타입의 제로값
}
