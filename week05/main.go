package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("iNPUT SCORE : ")
	reader := bufio.NewReader(os.Stdin)
	InputScore, err := reader.ReadString("\n") //err declared and not used
	fmt.Println(InputScore)
}
