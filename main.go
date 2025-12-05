package main

import (
	"log"

	"derekgarnett.com/aoc/2025/paperRolls"
)

func init() {
}

func main() {
	total := paperRolls.GetPaperRolls(true)

	log.Println(total)
}
