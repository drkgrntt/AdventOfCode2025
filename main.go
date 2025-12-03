package main

import (
	"log"

	"derekgarnett.com/aoc/2025/batteries"
)

func init() {
}

func main() {
	total := batteries.GetJoltage()

	log.Println(total)
}
