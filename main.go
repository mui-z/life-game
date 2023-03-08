package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	columnLength    = 10
	rowLength       = 10
	dieCellString   = "□"
	aliveCellString = "■"
)

func main() {

	var cells = generateAllDieCells()

	// glider
	cells[0][1] = aliveCellString
	cells[1][2] = aliveCellString
	cells[2][0] = aliveCellString
	cells[2][1] = aliveCellString
	cells[2][2] = aliveCellString

	// blink
	cells[0][7] = aliveCellString
	cells[1][7] = aliveCellString
	cells[2][7] = aliveCellString

	for {
		draw(cells)
		time.Sleep(500 * time.Millisecond)
		cells = calcNextGeneration(cells)
		clear()
	}
}

func calcNextGeneration(currentCells [][]string) [][]string {
	var nextCells = generateAllDieCells()

	for y := 0; y < columnLength; y++ {
		for x := 0; x < rowLength; x++ {
			var aliveCellCount = countAroundAliveCells(currentCells, x, y)
			var nextCell string
			currentCell := currentCells[y][x]

			if currentCell == dieCellString && aliveCellCount == 3 {
				nextCell = aliveCellString
			} else if currentCell == aliveCellString && (aliveCellCount == 2 || aliveCellCount == 3) {
				nextCell = aliveCellString
			} else {
				nextCell = dieCellString
			}

			nextCells[y][x] = nextCell
		}
	}

	return nextCells
}

func countAroundAliveCells(currentCells [][]string, x int, y int) int {
	var (
		upper  = 0
		middle = 0
		lower  = 0

		left  = 0
		right = 0
	)

	var leftOffset, rightOffset = calcHorizontalOffset(currentCells, x)

	if leftOffset != 0 && currentCells[y][x-leftOffset] == aliveCellString {
		left = 1
	}

	if rightOffset != 0 && currentCells[y][x+rightOffset] == aliveCellString {
		right = 1
	}

	if y != 0 {
		upper = strings.Count(strings.Join(currentCells[y-1][x-leftOffset:x+rightOffset+1], ""), aliveCellString)
	}

	if y != len(currentCells)-1 {
		lower = strings.Count(strings.Join(currentCells[y+1][x-leftOffset:x+rightOffset+1], ""), aliveCellString)
	}

	middle = left + right

	return upper + middle + lower
}

func calcHorizontalOffset(currentCells [][]string, x int) (int, int) {
	var (
		leftOffset  = 1
		rightOffset = 1
	)
	if x == 0 {
		leftOffset = 0
	}

	if x == len(currentCells[0])-1 {
		rightOffset = 0
	}

	return leftOffset, rightOffset
}

func generateAllDieCells() [][]string {
	cells := make([][]string, columnLength, columnLength)
	for i := 0; i < columnLength; i++ {
		cells[i] = make([]string, rowLength, rowLength)
	}

	for i := 0; i < columnLength; i++ {
		for j := 0; j < rowLength; j++ {
			cells[i][j] = dieCellString
		}
	}

	return cells
}

func draw(cells [][]string) {
	column := len(cells)
	row := len(cells[0])
	for i := 0; i < column; i++ {
		for j := 0; j < row; j++ {
			fmt.Printf(" " + cells[i][j])
		}
		fmt.Print("\n")
	}
}

func clear() {
	fmt.Print("\x1b[2J")
}
