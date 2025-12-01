package password

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var passwordInputFile = "password/input.txt"
var passwordInput = []string{}
var password = 0

func getInput() []string {
	if len(passwordInput) > 0 {
		return passwordInput
	}

	os.ReadFile(passwordInputFile)
	bytes, err := os.ReadFile(passwordInputFile)
	if err != nil {
		log.Fatal(err)
	}

	content := string(bytes)
	for _, line := range strings.Split(content, "\n") {
		if line != "" {
			passwordInput = append(passwordInput, line)
		}
	}

	return passwordInput
}

var currentPosition = 50

func handleInput(input string, useAlternateMethod bool) {
	dir := 0
	qty := 0
	var err error

	if strings.Contains(input, "R") {
		dir = 1
		pieces := strings.Split(input, "R")
		qty, err = strconv.Atoi(pieces[1])
	} else if strings.Contains(input, "L") {
		dir = -1
		pieces := strings.Split(input, "L")
		qty, err = strconv.Atoi(pieces[1])
	}

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < qty; i++ {
		currentPosition += (1 * dir)
		if currentPosition == 100 {
			currentPosition = 0
		}
		if currentPosition == -1 {
			currentPosition = 99
		}
		if currentPosition == 0 && useAlternateMethod {
			password += 1
		}
	}

	if currentPosition == 0 && !useAlternateMethod {
		password += 1
	}
}

func GetPassword(useAlternateMethod bool) int {
	for _, input := range getInput() {
		handleInput(input, useAlternateMethod)
	}

	return password
}
