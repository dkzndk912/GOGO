package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("iNPUT SCORE : ")
	reader := bufio.NewReader(os.Stdin)
	InputScoreString, err := reader.ReadString('\n') // Option 2
	if err != nil {
		log.Fatal(err)
	}

	InputScoreString = strings.TrimSpace(InputScoreString)      // remove space
	InputScore, err := strconv.ParseFloat(InputScoreString, 32) // casting

	if InputScore >= 90 {
		grade := "A grade!"
	} else {
		grade := "hmmm"
	}
	fmt.Println(grade)
}
