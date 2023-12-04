package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func RunTwo() {
	file, err := os.Open(FILE)
	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)
	}
	cards := cardsFromLines(lines)
	matchMap := createMatchMap(cards)
	cardIds := []int{}
	for _, card := range cards {
		cardIds = append(cardIds, cardIdsFromNumberOfMatches(card.id, matchMap)...)
		cardIds = append(cardIds, card.id)
	}
	sort.Ints(cardIds)
	fmt.Printf("Count: %d\n", len(cardIds))
}

func cardsFromLines(lines []string) []Card {
	cards := make([]Card, len(lines))
	for i, line := range lines {
		cards[i] = NewCard(line)
	}
	return cards
}

func createMatchMap(cards []Card) map[int]int {
	matchMap := map[int]int{}
	for _, card := range cards {
		matchMap[card.id] = len(card.matchingNums)
	}
	return matchMap
}

func cardIdsFromNumberOfMatches(sourceId int, matchMap map[int]int) []int {
	cardIds := []int{}
	for i := sourceId + 1; i < sourceId+1+matchMap[sourceId]; i++ {
		cardIds = append(cardIds, i)
	}
	newIds := []int{}
	for _, cardId := range cardIds {
		newIds = append(newIds, cardIdsFromNumberOfMatches(cardId, matchMap)...)
	}
	cardIds = append(cardIds, newIds...)
	return cardIds
}
