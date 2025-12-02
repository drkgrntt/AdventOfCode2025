package password

import (
	"log"
	"math"
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
var useAlternateMethod = false

func move(dir int, qty int) {
	if useAlternateMethod {
		password += int(math.Floor(float64(qty) / 100))
	}

	initialPosition := currentPosition
	movement := (qty % 100) * dir
	currentPosition += movement

	if currentPosition < 0 {
		currentPosition += 100
		// If we started at 0, we've already counted that
		if useAlternateMethod && initialPosition != 0 {
			password += 1
		}
	} else if currentPosition > 99 {
		currentPosition -= 100
		if useAlternateMethod {
			password += 1
		}
	} else if currentPosition == 0 {
		if useAlternateMethod {
			password += 1
		}
	}

	// If we're using the alternate method, this already happened
	if !useAlternateMethod && currentPosition == 0 {
		password += 1
	}
}

func parseInput(input string) (dir int, qty int) {
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

	return
}

func GetPassword(useAltMethod bool) int {
	useAlternateMethod = useAltMethod

	for _, input := range getInput() {
		dir, qty := parseInput(input)
		move(dir, qty)
	}

	return password
}
