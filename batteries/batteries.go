package batteries

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var batteryFile = "batteries/joltage.txt"
var batteryBanks = []string{
	// // answer is 357
	// "987654321111111",
	// "811111111111119",
	// "234234234234278",
	// "818181911112111",
}
var total = 0

func getBatteries() []string {
	if len(batteryBanks) > 0 {
		return batteryBanks
	}

	bytes, err := os.ReadFile(batteryFile)
	if err != nil {
		log.Fatal(err)
	}

	content := string(bytes)
	for _, line := range strings.Split(content, "\n") {
		if line != "" {
			batteryBanks = append(batteryBanks, line)
		}
	}

	return batteryBanks
}

func handleBank(bank string, numDigits int) {
	positions := []int{}
	digits := []string{}

	// We're running the process to find the digit once per number of digits
	// Starting at the end so we know how far from the end of the string we need to stop
digitsLoop:
	for x := numDigits; x > 0; x-- {
		var position int
		if len(positions) == 0 {
			position = -1
		} else {
			position = positions[len(positions)-1]
		}

		// Loop backwards from 9-1 to prioritize the largest value
	optionsLoop:
		for i := 9; i > 0; i-- {

			// Go through each batter in the battery bank
			for j := position + 1; j < len(bank); j++ {

				// If we've reached the end of the battery bank while leaving the correct number of digits available,
				// then this value was not available for this digit
				if j == len(bank)-x+1 {
					continue optionsLoop
				}

				// If we've found the digit that matches, add it to the list
				if string(bank[j]) == fmt.Sprint(i) {
					positions = append(positions, j)
					digits = append(digits, string(bank[j]))

					// Move on to the next digit
					continue digitsLoop
				}
			}
		}
	}

	var joltageString string
	for _, digit := range digits {
		joltageString += digit
	}

	joltage, err := strconv.Atoi(joltageString)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(joltage)

	if err != nil {
		log.Fatal(err)
	}

	total += joltage
}

func GetJoltage(numDigits int) int {
	for _, bank := range getBatteries() {
		handleBank(bank, numDigits)
	}
	return total
}
