package day03

import "testing"

func TestPart1Example(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := solve1(test.input)

			if actual != test.expected {
				t.Errorf("Part1(%s) = %d, expected %d", test.input, actual, test.expected)
			}
		})
	}
}

func TestPart1Solution(t *testing.T) {
	expected := 2081
	actual := Part1()

	if actual != expected {
		t.Errorf("Part1() = %d, expected %d", actual, expected)
	}
}

func TestPart2Example(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := solve1(test.input)

			if actual != test.expected {
				t.Errorf("Part1(%s) = %d, expected %d", test.input, actual, test.expected)
			}
		})
	}
}

func TestPart2Solution(t *testing.T) {
	expected := 2341
	actual := Part2()

	if actual != expected {
		t.Errorf("Part2() = %d, expected %d", actual, expected)
	}
}
