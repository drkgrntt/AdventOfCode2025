package main

import (
	"log"

	"derekgarnett.com/aoc/2025/password"
)

func init() {
}

func main() {
	useAlternateMethod := true
	pass := password.GetPassword(useAlternateMethod)

	log.Println("The password is: ", pass)
}
