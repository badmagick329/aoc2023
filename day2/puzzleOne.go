package day2

import (
	"fmt"
)

func RunOne() {
	bag := Colors{
		12,
		13,
		14,
	}
	lines := linesFromFile()
	gameIds := possibleGameIds(lines, bag)
	sum := 0
	for gameId := range gameIds {
		sum += gameId
	}
	fmt.Printf("Sum: %d\n", sum)
}

func possibleGameIds(lines chan string, bag Colors) chan int {
	gameIds := make(chan int)
	gameNum := 1
	go func() {
		defer close(gameIds)
		for line := range lines {
			pulls := parseColors(line)
			impossible := false
			for _, pull := range pulls {
				impossible = !bag.contains(pull)
				if impossible {
					break
				}
			}
			if !impossible {
				gameIds <- gameNum
			}
			gameNum++
		}
	}()

	return gameIds
}
