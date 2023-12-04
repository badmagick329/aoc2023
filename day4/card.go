package day4

import (
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	id           int
	winningNums  []int
	myNums       []int
	matchingNums []int
}

func NewCard(line string) Card {
	winningNums, myNums, id := readNumsAndId(line)
	card := Card{
		id,
		winningNums,
		myNums,
		[]int{},
	}
	card.calculateMatches()
	return card
}

func readNumsAndId(line string) ([]int, []int, int) {
	idAndNums := strings.Split(line, ": ")
	pattern := regexp.MustCompile(`Card\s+(\d+)`)
	idStr := pattern.FindStringSubmatch(idAndNums[0])[1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatalf("Error converting %s", idStr)
	}
	numsStr := strings.Split(idAndNums[1], " | ")
	winningNumsStr, myNumsStr := numsStr[0], numsStr[1]
	return strToNums(winningNumsStr), strToNums(myNumsStr), id
}

func strToNums(numsStr string) []int {
	nums := []int{}
	for _, ns := range strings.Split(numsStr, " ") {
		conv, err := strconv.Atoi(ns)
		if err != nil {
			continue
		}
		nums = append(nums, conv)
	}
	return nums
}

func (c *Card) calculateMatches() {
	for _, mn := range c.myNums {
		if slices.Contains(c.winningNums, mn) {
			c.matchingNums = append(c.matchingNums, mn)
		}
	}
}

func (c *Card) CalculatePoints() int {
	if len(c.matchingNums) == 0 {
		return 0
	}
	points := 1
	for i := 1; i < len(c.matchingNums); i++ {
		points *= 2
	}
	return points
}
