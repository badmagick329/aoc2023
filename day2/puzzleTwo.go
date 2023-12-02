package day2

import (
	"fmt"
)

func RunTwo() {
	lines := linesFromFile()
	minColors := minColorsRequired(lines)
	sum := 0
	for mc := range minColors {
		sum += mc.power()
	}
	fmt.Printf("Sum: %d\n", sum)
}

func minColorsRequired(lines chan string) chan Colors {
	minColors := make(chan Colors)
	go func() {
		defer close(minColors)
		for line := range lines {
			pulls := parseColors(line)
			minColors <- minRequiredFor(pulls)
		}
	}()
	return minColors
}
