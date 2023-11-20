package main

import "fmt"

func main() {
	a := []string{"a", "b", " c", "d"}
	a = append(a, "e")
	as := a[:2]
	as[1] = "2"
	fmt.Println(a)
	fmt.Println(as)

	b := make([]string, 4, 5)
	b[0] = "a"
	b[2] = "c"
	b[3] = "d"
	bs := b[0:2]
	b[0] = "a"

	fmt.Println(b, len(b), cap(b))

}
