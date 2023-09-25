package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	//var a int
	//a = 9

	//var a = 9

	var a int = 9

	//a := 9    // 숏폼
	b := 2.71 // float64

	fmt.Println(float64(a) * b)
	//d, e := 4.10, 3.99

	fmt.Println(a)
	fmt.Println(reflect.TypeOf('z')) ///rune. int 32
	fmt.Println(reflect.TypeOf(99))
	fmt.Println(reflect.TypeOf(2.7))
	fmt.Println(reflect.TypeOf(false))
	fmt.Println(reflect.TypeOf("Go!"))

	fmt.Println('A', 'a', '0', '김')
	fmt.Println(math.Floor(3.17))
	fmt.Println(math.Ceil(3.17))
	// fmt.Println(strings.Title("오픈소스프로그래밍~"))
	fmt.Println(strings.Title("open source programming~"))
}
