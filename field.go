package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func getField(size, maxLenShip int, verbose bool) [][]int {
	field := makeEmptyField(size)
	for len := maxLenShip; len >= 1; len-- {
		times := maxLenShip - len + 1
		for i := 1; i <= times; i++ {
			putShip(len, field)
			if verbose {
				fmt.Printf("Ship of len %d: #%d\n", len, i)
				printField(field)
			}
		}
	}
	return field
}

func makeEmptyField(size int) [][]int {
	field := make([][]int, size)
	for row := range field {
		field[row] = make([]int, size)
	}
	return field
}

func printField(field [][]int) {
	for row := range field {
		fmt.Println(
			strings.Trim(
				strings.Join(
					strings.Fields(
						fmt.Sprint(field[row])), "  "), "[]"))
	}
}

func putShip(lenShip int, field [][]int) [][]int {
	dir := rand.Intn(2) == 1
	availableBegs := getAvailableBegs(lenShip, field, dir)
	choice := availableBegs[rand.Intn(len(availableBegs))]
	begRow := choice[0]
	begCol := choice[1]
	if dir {  // downwards
		for row := begRow; row < begRow+lenShip; row++ {
			field[row][begCol] = 1
		}
	} else {  // to the right
		for col := begCol; col < begCol+lenShip; col++ {
			field[begRow][col] = 1
		}
	}
	return field
}

func getAvailableBegs(lenShip int, field [][]int, revert bool) [][]int {
	availableBegs := make([][]int, 0)
	size := len(field)
	var currField [][]int
	if revert {
		currField = revertField(field)
	} else {
		currField = field
	}
	for col := 0; col <= size-lenShip; col++ {
		for row := range currField {
			var begCoord []int
			if revert {
				begCoord = append(begCoord, col, row)
			} else {
				begCoord = append(begCoord, row, col)
			}
			isSuitable := true
			for step := col; (step <= col+lenShip) && (step < size); step++ {
				if isUnsuitable(row, step, currField) {
					isSuitable = false
					break
				}
			}
			if isSuitable {
				availableBegs = append(availableBegs, begCoord)
			}
		}
	}
	return availableBegs
}

func revertField(field [][]int) [][]int {
	revertedField := makeEmptyField(len(field))
	for row := range field {
		for col := range field {
			revertedField[row][col] = field[col][row]
		}
	}
	return revertedField
}

func isUnsuitable(row, step int, currField [][]int) bool {
	size := len(currField)
	return (currField[row][step] != 0) || 
	((row - 1 >= 0) && currField[row-1][step] != 0) || 
	((row + 1 < size) && currField[row+1][step] != 0) 
}
