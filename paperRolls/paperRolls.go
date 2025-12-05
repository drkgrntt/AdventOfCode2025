package paperRolls

import (
	"log"
	"os"
	"strings"
)

var gridFile = "paperRolls/grid.txt"
var gridRows = []string{
	// "..@@.@@@@.",
	// "@@@.@.@.@@",
	// "@@@@@.@.@@",
	// "@.@@@@..@.",
	// "@@.@@@@.@@",
	// ".@@@@@@@.@",
	// ".@.@.@.@@@",
	// "@.@@@.@@@@",
	// ".@@@@@@@@.",
	// "@.@.@@@.@.",
}
var grid = [][]string{}
var total = 0

func getGrid() [][]string {
	if len(grid) > 0 {
		return grid
	}

	if len(gridRows) == 0 {
		bytes, err := os.ReadFile(gridFile)
		if err != nil {
			log.Fatal(err)
		}

		content := string(bytes)
		for _, line := range strings.Split(content, "\n") {
			if line != "" {
				gridRows = append(gridRows, line)
			}
		}
	}

	for _, row := range gridRows {
		grid = append(grid, strings.Split(row, ""))
	}

	return grid
}

func checkRoll(row int, col int) {
	adjacentRolls := 0
	grid := getGrid()

	isFirstRow := row == 0
	isLastRow := row == len(grid)-1
	isFirstCol := col == 0
	isLastCol := col == len(grid[0])-1

	// ooo
	// x@x
	// xxx
	if !isFirstRow {
		if !isFirstCol {
			if grid[row-1][col-1] == "@" {
				adjacentRolls++
			}
		}

		if grid[row-1][col] == "@" {
			adjacentRolls++
		}

		if !isLastCol {
			if grid[row-1][col+1] == "@" {
				adjacentRolls++
			}
		}
	}

	// xxx
	// o@x
	// oxx
	if !isFirstCol {
		if grid[row][col-1] == "@" {
			adjacentRolls++
		}

		if !isLastRow {
			if grid[row+1][col-1] == "@" {
				adjacentRolls++
			}
		}
	}

	// xxx
	// x@x
	// xoo
	if !isLastRow {
		if grid[row+1][col] == "@" {
			adjacentRolls++
		}

		if !isLastCol {
			if grid[row+1][col+1] == "@" {
				adjacentRolls++
			}
		}
	}

	// xxx
	// x@o
	// xxx
	if !isLastCol {
		if grid[row][col+1] == "@" {
			adjacentRolls++
		}
	}

	if adjacentRolls > 3 {
		return
	}

	log.Println(row, col)
	total++
}

func GetPaperRolls() int {
	// log.Println(getGrid())
	for rowI, row := range getGrid() {
		for colI, col := range row {
			if col == "@" {
				checkRoll(rowI, colI)
			}
		}
	}
	return total
}
