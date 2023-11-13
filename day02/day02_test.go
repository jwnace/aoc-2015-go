package day02

import "testing"

func TestGetWrappingPaperRequired(t *testing.T) {
	tests := []struct {
		name     string
		l        int
		w        int
		h        int
		expected int
	}{
		{"2x3x4", 2, 3, 4, 58},
		{"1x1x10", 1, 1, 10, 43},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := getWrappingPaperRequired(test.l, test.w, test.h)

			if actual != test.expected {
				t.Errorf("Part1(%s) = %d, expected %d", test.name, actual, test.expected)
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

func TestGetRibbonRequired(t *testing.T) {
	tests := []struct {
		name     string
		l        int
		w        int
		h        int
		expected int
	}{
		{"2x3x4", 2, 3, 4, 34},
		{"1x1x10", 1, 1, 10, 14},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := getRibbonRequired(test.l, test.w, test.h)

			if actual != test.expected {
				t.Errorf("Part1(%s) = %d, expected %d", test.name, actual, test.expected)
			}
		})
	}
}

func TestPart2Solution(t *testing.T) {
	expected := 3737498
	actual := Part2()

	if actual != expected {
		t.Errorf("Part2() = %d, expected %d", actual, expected)
	}
}
