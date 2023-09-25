package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("iNPUT SCORE : ")
	reader := bufio.NewReader(os.Stdin)
	InputScore, _ := reader.ReadString('\n') // Option 1, ignore the errot return value with the blank identifier
	fmt.Println(InputScore)
}
