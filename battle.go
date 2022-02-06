package main

import (
	// "fmt"
	"math/rand"
)

func main() {
	rand.Seed(4)
	field := getField(10, 4, false)
	printField(field)
	// field := getTestField()
	// printField(field)
	// availableBegs := getAvailableBegs(4, field, false)
	// fmt.Println(availableBegs)
	// revertedField := revertField(field)
	// fmt.Println()
	// printField(revertedField)
}