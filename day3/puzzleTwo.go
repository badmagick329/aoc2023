package day3

import (
	"fmt"
	"os"
	"strings"
)

func RunTwo() {
	dat, _ := os.ReadFile(FILE)
	eng := NewEngine(strings.Split(string(dat), "\n"))
	ratios := []int{}
	for _, gc := range eng.gearCells() {
		r := gc.gearRatio()
		// fmt.Printf("Ratio: %d\n", r)
		ratios = append(ratios, r)
	}
	sum := 0
	for _, n := range ratios {
		sum += n
	}
	fmt.Printf("Sum: %d\n", sum)
}
