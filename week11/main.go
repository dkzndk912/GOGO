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

	i := 0
	for i < len(primes) {
		fmt.Println(primes[i])
		i++
	}

	for idx, prime := range primes { // idx를 빼서 값만 출력하라 했으나 인덱스가 출력됨
		fmt.Println(idx, prime) // 선언된 변수를 사용하지 않으면 에러난다
	}
	for _, prime := range primes { // 빈식별자로 해결
		fmt.Println(prime)
	}
	fmt.Printf("%#v", primes)
	fmt.Printf("%#v", test)
}
