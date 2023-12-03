package main

import (
	"regexp"
	"strconv"
	s "strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

var alphabetRegex = regexp.MustCompile(`[^0-9]+`)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	inputArr := s.Split(input, "\n")
	sumCalVals := 0
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	for _, str := range inputArr {
		substr := alphabetRegex.ReplaceAllString(str, "")
		val, _ := strconv.Atoi(string(substr[0]) + string(substr[len(substr)-1]))
		sumCalVals += val
	}
	return sumCalVals
}