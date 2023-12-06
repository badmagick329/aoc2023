package day6

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"os"
)

func RunOne() {
	file, _ := os.Open(FILE)
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	times := lib.ParseNumsFromText(lines[0])
	distances := lib.ParseNumsFromText(lines[1])
	waysToBeat := []int{}
	for i := 0; i < len(times); i++ {
		waysToBeat = append(waysToBeat, nWaysToBeat(times[i], distances[i]))
	}
	fmt.Println(waysToBeat)
	prod := 1
	for _, n := range waysToBeat {
		prod *= n
	}
	fmt.Println(prod)
}
