package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func RunOne() {
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
	mapper := NewMapper(lines)
	locationNums := []int{}
	// seeds := []int{79, 14, 55, 13}
	for _, seed := range mapper.seeds {
		fmt.Printf("Getting location for seed: %d\n", seed)
		num, err := mapper.Transform("seed", "location", seed)
		if err != nil {
			log.Fatal(err)
		}
		locationNums = append(locationNums, num)
	}
	sort.Ints(locationNums)
	fmt.Println(locationNums)
}
