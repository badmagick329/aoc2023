package day3

import (
	"fmt"
	"os"
	"strings"
)

func RunOne() {
	dat, _ := os.ReadFile(FILE)
	eng := NewEngine(strings.Split(string(dat), "\n"))
	sum := 0
	for _, n := range eng.PartsAsNumbers() {
		sum += n
	}
	fmt.Printf("Sum: %d\n", sum)
}
