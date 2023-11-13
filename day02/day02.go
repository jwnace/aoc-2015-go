package day02

import (
	"joenace.com/aoc-2015-go/helpers"
	"strconv"
	"strings"
)

func Part1() int {
	input := helpers.ReadInput()
	return Solve1(input)
}

func Part2() int {
	input := helpers.ReadInput()
	return Solve2(input)
}

func Solve1(input string) int {
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

		sum += getWrappingPaperRequired(l, w, h)
	}

	return sum
}

func Solve2(input string) int {
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

		sum += getRibbonRequired(l, w, h)
	}

	return sum
}

func getWrappingPaperRequired(l int, w int, h int) int {
	return 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
}

func getRibbonRequired(l int, w int, h int) int {
	return 2*min(l+w, w+h, h+l) + l*w*h
}
