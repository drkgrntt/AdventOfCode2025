package main

import (
	"log"

	"derekgarnett.com/aoc/2025/invalidIds"
)

func init() {
}

func main() {
	total := invalidIds.GetInvalidIdTotal(true)

	log.Println(total)
}
