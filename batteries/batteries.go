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

func handleBank(bank string) {
	firstPosition := -1
	var firstDigit, secondDigit string

digits:
	for x := 0; x < 2; x++ {
	gtl:
		for i := 9; i > 0; i-- {
			for j := firstPosition + 1; j < len(bank); j++ {
				if firstPosition == -1 && j == len(bank)-1 {
					continue gtl
				}
				if string(bank[j]) == fmt.Sprint(i) {
					if firstPosition == -1 {
						firstPosition = j
						firstDigit = string(bank[j])
					} else {
						secondDigit = string(bank[j])
					}
					continue digits
				}
			}
		}
	}

	joltage, err := strconv.Atoi(firstDigit + secondDigit)

	// log.Println(joltage)

	if err != nil {
		log.Fatal(err)
	}

	total += joltage
}

func GetJoltage() int {
	for _, bank := range getBatteries() {
		handleBank(bank)
	}
	return total
}
