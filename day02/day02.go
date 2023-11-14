package day02

import (
	"joenace.com/aoc-2015-go/helpers"
	"strconv"
	"strings"
)

var inputFile string

func init() {
	inputFile = helpers.ReadInput()
}

func Part1() int {
	return solve1(inputFile)
}

func Part2() int {
	return solve2(inputFile)
}

func solve1(input string) int {
	return solve(input, getWrappingPaperRequired)
}

func solve2(input string) int {
	return solve(input, getRibbonRequired)
}

func solve(input string, calculate func(int, int, int) int) int {
	sum := 0

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.Split(line, "x")

		l, err := strconv.Atoi(parts[0])

		if err != nil {
			panic(err)
		}

		w, err := strconv.Atoi(parts[1])

		if err != nil {
			panic(err)
		}

		h, err := strconv.Atoi(parts[2])

		if err != nil {
			panic(err)
		}

		sum += calculate(l, w, h)
	}

	return sum
}

func getWrappingPaperRequired(l int, w int, h int) int {
	return 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
}

func getRibbonRequired(l int, w int, h int) int {
	return 2*min(l+w, w+h, h+l) + l*w*h
}
