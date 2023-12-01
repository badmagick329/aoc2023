package day1

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const FILE = "day1/file.txt"

var NUMS = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var NUM_MAP = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func Run() {
	dat, err := os.ReadFile(FILE)
	if err != nil {
		log.Fatalf("Error reading file %s", FILE)
	}
	contents := string(dat)
	sum := _run(contents)
	fmt.Printf("Answer: %d\n", sum)
}

func _run(input string) int {
	samples := strings.Fields(input)
	nums := make([]int, len(samples))
	sum := 0
	for i, sample := range samples {
		if strings.TrimSpace(sample) == "" {
			continue
		}
		nums[i] = getNums(sample)
		sum += nums[i]
	}
	return sum
}

func getNums(str string) int {
	first, last := -1, -1
	runes := []rune(str)
	num, err := findFirst(runes)
	if err == nil {
		first, last = num, num
	}
	num, err = findLast(runes)
	if err == nil {
		last = num
		if first == -1 {
			first = num
		}
	}
	if first == -1 || last == -1 {
		log.Fatalf("Couldn't find first or last number in %s\n", str)
	}
	strNum := fmt.Sprintf("%d%d", first, last)
	num, err = strconv.Atoi(strNum)
	if err != nil {
		log.Fatal("Huh")
	}
	return num
}

func findFirst(runes []rune) (int, error) {
	max := len(runes)
	endIdx := 1
	for endIdx < max {
		num, err := getNum(string(runes[:endIdx]))
		endIdx++
		if err != nil {
			continue
		}
		return num, nil
	}
	return 0, fmt.Errorf("No number found in %v\n", string(runes))
}

func findLast(runes []rune) (int, error) {
	startIdx := len(runes) - 1
	for startIdx > 0 {
		num, err := getNum(string(runes[startIdx:]))
		startIdx--
		if err != nil {
			continue
		}
		return num, nil
	}
	return 0, fmt.Errorf("No number found in %v\n", string(runes))
}

func getNum(str string) (int, error) {
	for _, s := range NUMS {
		if strings.Contains(str, s) {
			num, err := strconv.Atoi(s)
			if err != nil {
				return 0, err
			}
			return num, nil
		}
	}
	for k, v := range NUM_MAP {
		if strings.Contains(str, k) {
			return v, nil
		}
	}
	return 0, fmt.Errorf("No num found in %s\n", str)
}
