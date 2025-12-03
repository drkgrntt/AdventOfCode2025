package main

import (
	"log"

	"derekgarnett.com/aoc/2025/batteries"
)

func init() {
}

func main() {
	total := batteries.GetJoltage(12)

	log.Println(total)
}
