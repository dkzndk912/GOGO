package mymath

func MyAbs(num int) int { // 외부 파일에 함수를 노출하려면 첫글자를 대문자로 내부만사용 소문자
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
