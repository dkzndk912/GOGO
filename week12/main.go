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

	println("----------------")

	copyS := s[1:4]

	for _, v := range copyS {
		fmt.Println(v)
	}

	println("----------------")

	test := [3]string{"inha", "go", "student"} // test 배열생성
	// testS := test[0:2]
	// fmt.Println(len(testS))

	testS := test[:2]
	test2S := test[1:]
	test2S[0] = "python"
	fmt.Println(test2S[1])
	fmt.Println(testS, len(testS))
}
