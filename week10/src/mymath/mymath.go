package mymath

func MyAbs(num int) int {
	if num > -1 {
		return num
	} else {
		return num * -1
	}
}

func MyPower(base int, exponent int) int {
	result := 1
	for i := exponent; i >= 1; i-- {
		result *= base
	}
	return result
}
