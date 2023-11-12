package day02

import "testing"

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
