package day3

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Cell struct {
	pos    Pos
	val    string
	isPart bool
}

type GearCell struct {
	pos         Pos
	partNumbers []PartNumber
}

func (gc *GearCell) gearRatio() int {
	return gc.partNumbers[0].val * gc.partNumbers[1].val
}

func (c *Cell) String() string {
	str := strings.Builder{}
	str.WriteString(
		fmt.Sprintf(
			"[%d, %d (%s)",
			c.pos.x,
			c.pos.y,
			c.val,
		))
	if c.isPart {
		str.WriteString("P")
	} else {

		str.WriteString(" ")
	}
	if c.isSymbol() {
		str.WriteString("S")
	} else {
		str.WriteString(" ")
	}
	str.WriteString("]")
	return str.String()
}

func (c *Cell) isSymbol() bool {
	return !c.isDigit() && c.val != "."
}

func (c *Cell) isDigit() bool {
	return unicode.IsDigit([]rune(c.val)[0])
}

func ToPartNumbers(cells []Cell) []PartNumber {
	partPositions := []Pos{}
	partNumbers := []PartNumber{}
	str := strings.Builder{}
	for _, cell := range cells {
		if !cell.isPart && len(partPositions) != 0 {
			numVal, err := strconv.Atoi(str.String())
			if err != nil {
				log.Fatalf("Error convering %s\n", str.String())
			}
			partNumbers = append(partNumbers, NewPartNumber(partPositions, numVal))
			str.Reset()
			partPositions = []Pos{}
			continue
		}
		if cell.isPart {
			partPositions = append(partPositions, cell.pos)
			str.WriteString(cell.val)
		}
	}
	if len(partPositions) != 0 {
		conv, err := strconv.Atoi(str.String())
		if err != nil {
			log.Fatalf("Error convering %s\n", str.String())
		}
		partNumbers = append(partNumbers, NewPartNumber(partPositions, conv))
	}
	return partNumbers
}
func _ToPartNumbers(cells []Cell) []int {
	nums := []int{}
	numStrings := []string{}
	str := strings.Builder{}
	for _, cell := range cells {
		if !cell.isPart && str.Len() != 0 {
			numStrings = append(numStrings, str.String())
			str.Reset()
			continue
		}
		if cell.isPart {
			str.WriteString(cell.val)
		}
	}
	if str.Len() != 0 {
		numStrings = append(numStrings, str.String())
	}
	for _, ns := range numStrings {
		conv, err := strconv.Atoi(ns)
		if err != nil {
			log.Fatalf("Error convering %s\n", ns)
		}
		nums = append(nums, conv)
	}
	return nums
}
