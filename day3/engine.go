package day3

import (
	"strings"
)

type Engine struct {
	cells       [][]Cell
	partNumbers []PartNumber
}

func (e *Engine) String() string {
	str := strings.Builder{}
	for _, cy := range e.cells {
		for _, cx := range cy {
			str.WriteString(cx.String())
			str.WriteString(" ")
		}
		str.WriteString("\n\n")
	}
	return str.String()
}

func NewEngine(lines []string) Engine {
	eng := Engine{
		cells: linesToCells(lines),
	}
	eng.markParts()
	eng.initPartNumbers()
	return eng
}

func linesToCells(lines []string) [][]Cell {
	cellsArr := [][]Cell{}
	for y, line := range lines {
		cells := []Cell{}
		for x, ch := range line {
			cells = append(cells, Cell{
				pos: Pos{x, y},
				val: string(ch),
			})
		}
		if len(cells) > 0 {
			cellsArr = append(cellsArr, cells)
		}
	}
	return cellsArr
}

func (e *Engine) markParts() {
	for _, cy := range e.cells {
		for _, cx := range cy {
			if cx.isSymbol() {
				e.markAdjacent(cx.pos)
			}
		}
	}
}

func (e *Engine) at(p Pos) *Cell {
	return &e.cells[p.y][p.x]
}

func (e *Engine) markAdjacent(pos Pos) {
	adjacents := []Pos{}
	markedPositions := []Pos{}
	for _, p := range pos.adjacentPositions(e.width(), e.height()) {
		if e.at(p).val != "." {
			adjacents = append(adjacents, p)
		}
	}
	for _, p := range adjacents {
		if e.at(p).isDigit() && !e.at(p).isPart {
			e.at(p).isPart = true
			markedPositions = append(markedPositions, p)
		}
	}
	for _, p := range markedPositions {
		e.markAdjacent(p)
	}
}

func (e *Engine) width() int {
	return len(e.cells[0])
}

func (e *Engine) height() int {
	return len(e.cells)
}

func (e *Engine) PartsAsNumbers() []int {
	nums := []int{}
	for _, pn := range e.partNumbers {
		nums = append(nums, pn.val)
	}
	return nums
}

func (e *Engine) initPartNumbers() {
	for _, cells := range e.cells {
		pnums := ToPartNumbers(cells)
		for _, pn := range pnums {
			e.partNumbers = append(e.partNumbers, pn)
		}
	}
}

func (e *Engine) adjacentPartNumbers(pos Pos) []PartNumber {
	adjacents := pos.adjacentPositions(e.width(), e.height())
	partNumbers := []PartNumber{}
	for _, pn := range e.partNumbers {
		if pn.InPositions(adjacents) {
			partNumbers = append(partNumbers, pn)
		}
	}
	partNumbers = UniquePartNumbers(partNumbers)
	return partNumbers
}

func (e *Engine) gearCells() []GearCell {
	gc := []GearCell{}
	for _, cy := range e.cells {
		for _, cx := range cy {
			if cx.val == "*" {
				partNumbers := e.adjacentPartNumbers(cx.pos)
				if len(partNumbers) != 2 {
					continue
				}
				gc = append(gc, GearCell{
					cx.pos,
					partNumbers,
				})
			}
		}
	}
	return gc
}
