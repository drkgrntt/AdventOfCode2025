package invalidIds

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var rangesFile = "invalidIds/ranges.txt"
var ranges = []string{}
var total = 0
var useAlternateMethod = false

func getRanges() []string {
	if len(ranges) > 0 {
		return ranges
	}

	os.ReadFile(rangesFile)
	bytes, err := os.ReadFile(rangesFile)
	if err != nil {
		log.Fatal(err)
	}

	content := string(bytes)
	content = strings.ReplaceAll(content, "\n", "")

	for _, line := range strings.Split(content, ",") {
		if line != "" {
			ranges = append(ranges, line)
		}
	}

	return ranges
}

type Range struct {
	start int
	end   int
}

func parseRange(idRange string) *Range {
	pieces := strings.Split(idRange, "-")

	start, err := strconv.Atoi(pieces[0])
	end, err := strconv.Atoi(pieces[1])

	if err != nil {
		log.Println(err)
	}
	return &Range{
		start: start,
		end:   end,
	}
}

func handleRange(idRange *Range) {
	for id := idRange.start; id <= idRange.end; id++ {
		stringId := strconv.Itoa(id)
		size := len(stringId)
		halfSize := size / 2

		if !useAlternateMethod {
			firstHalf := stringId[:halfSize]
			secondHalf := stringId[halfSize:]

			if firstHalf == secondHalf {
				total += id
			}
		} else {

			for i := 1; i <= halfSize; i++ {
				firstPiece := stringId[:i]
				remainder := strings.ReplaceAll(stringId, firstPiece, "")

				if remainder == "" {
					total += id
					break
				}
			}
		}
	}
}

func GetInvalidIdTotal(altMethod bool) int {
	useAlternateMethod = altMethod
	for _, input := range getRanges() {
		idRange := parseRange(input)
		handleRange(idRange)
	}

	return total
}
