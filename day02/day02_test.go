package day02

import "testing"

func TestPart1Example(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		// TODO: add test cases
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := Solve1(test.input)

			if actual != test.expected {
				t.Errorf("Part1(%s) = %d, expected %d", test.input, actual, test.expected)
			}
		})
	}
}

func TestPart1Solution(t *testing.T) {
	expected := 1586300
	actual := Part1()

	if actual != expected {
		t.Errorf("Part1() = %d, expected %d", actual, expected)
	}
}

func TestPart2Solution(t *testing.T) {
	expected := 3737498
	actual := Part2()

	if actual != expected {
		t.Errorf("Part2() = %d, expected %d", actual, expected)
	}
}
