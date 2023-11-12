package day01

import "joenace.com/aoc-2015-go/helpers"

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

	for _, value := range input {
		if value == '(' {
			sum++
		} else if value == ')' {
			sum--
		}
	}

	return sum
}

func Solve2(input string) int {
	sum := 0

	for i, value := range input {
		if value == '(' {
			sum++
		} else if value == ')' {
			sum--

			if sum < 0 {
				return i + 1
			}
		}
	}

	return 0
}
