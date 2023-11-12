package day01

import (
	"os"
	"path"
	"runtime"
)

func loadInput() string {
	_, file, _, _ := runtime.Caller(0)
	dir := path.Dir(file)
	input, err := os.ReadFile(dir + "/input.txt")

	if err != nil {
		return ""
	}

	return string(input)
}

func Part1() int {
	input := loadInput()
	return Solve1(input)
}

func Part2() int {
	input := loadInput()
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
