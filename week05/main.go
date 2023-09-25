package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // get ther current datat and time as an intger
	answer := rand.Intn(100) + 1 // random intger number betwwen
	fmt.Println("game start")
	fmt.Println((answer))

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		fmt.Println("You have", 10-i, "chances~")
		fmt.Print("Input guesss number : ")
		inputNumberString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		inputNumberString = strings.TrimSpace(inputNumberString)
		inputNumber, err := strconv.Atoi(inputNumberString)

		if err != nil {
			log.Fatal(err)
		}

		if inputNumber < answer {
			fmt.Println("Guess number is lower than answer")
		} else if inputNumber > answer {
			fmt.Println("Guess number is higher than answer")
		} else {
			fmt.Println("RIGHT!")
			break

		}
	}
}
