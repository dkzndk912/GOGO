package week05

import (
	"fmt"
	"math/rand"
)

func main() {

	dice := rand.Intn(6) + 1
	fmt.Println((dice))
}
