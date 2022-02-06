package main

import (
	// "fmt"
	"math/rand"
)

func main() {
	rand.Seed(4)
	field := getField(10, 4, false)
	printField(field)
}
