package day6

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunTwo() {
	file, _ := os.Open(FILE)
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	time := parseNum(lines[0])
	distance := parseNum(lines[1])
	waysToBeat := nWaysToBeat(time, distance)
	fmt.Println(waysToBeat)
}

func parseNum(line string) int {
	nums := lib.ParseNumsFromText(line)
	str := strings.Builder{}
	for _, n := range nums {
		str.WriteString(fmt.Sprintf("%d", n))
	}
	conv, _ := strconv.Atoi(str.String())
	return conv
}
