package day3

import (
	"log"
	"slices"
)

type PartNumber struct {
	id  Pos
	pos []Pos
	val int
}

func NewPartNumber(pos []Pos, val int) PartNumber {
	if len(pos) < 1 {
		log.Fatal("Cannot create part number from an empty pos array")
	}
	pn := PartNumber{
		id:  pos[0],
		pos: pos,
		val: val,
	}
	return pn
}

func (pn *PartNumber) InPositions(positions []Pos) bool {
	for _, pos := range positions {
		if slices.Contains(pn.pos, pos) {
			return true
		}
	}
	return false
}

func (pn *PartNumber) InPartNumbers(partNumbers []PartNumber) bool {
	for _, p := range partNumbers {
		if pn.id == p.id {
			return true
		}
	}
	return false
}

func UniquePartNumbers(partNumbers []PartNumber) []PartNumber {
	uniques := []PartNumber{}
	for _, pn := range partNumbers {
		if !pn.InPartNumbers(uniques) {
			uniques = append(uniques, pn)
		}
	}
	return uniques
}
