package main

import (
	"fmt"
	"math/rand"
)

func main() {
	message := "My favorite number is"

	fmt.Println(message, rand.Intn(1000))
}