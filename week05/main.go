package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("iNPUT SCORE : ")
	reader := bufio.NewReader(os.Stdin)
	InputScore, err := reader.ReadString('\n') // Option 2
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(InputScore)
}
