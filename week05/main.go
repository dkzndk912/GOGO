package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("iNPUT SCORE : ")
	reader := bufio.NewReader(os.Stdin)
	InputScore := reader.ReadString("\n") // 1 variable but reader.ReadString returns 2 values
	fmt.Println(InputScore)
}
