package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
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
	mapper := NewMapper2(lines)
	minsChan := make(chan int, len(mapper.seedRanges))
	processedSeedsCount := make(chan int, len(mapper.seedRanges))
	var wg sync.WaitGroup
	// seeds := []int{79, 14, 55, 13}
	for i, seedRange := range mapper.seedRanges {
		fmt.Printf("[%d/%d] Processing seed range: %v\n", i+1, len(mapper.seedRanges), seedRange)
		wg.Add(1)
		go func(seedRange SeedRange) {
			defer wg.Done()
			minNum := -1
			count := 0
			for i := seedRange.start; i < seedRange.start+seedRange.rangeLen; i++ {
				res := process_seed(mapper, i)
				if minNum == -1 || res < minNum {
					minNum = res
				}
				count++
			}
			minsChan <- minNum
			processedSeedsCount <- count
		}(seedRange)
	}
	wg.Wait()
	close(minsChan)
	close(processedSeedsCount)
	minNum := -1
	for n := range minsChan {
		if minNum == -1 || n < minNum {
			minNum = n
		}
	}
	fmt.Printf("Min num is: %d\n", minNum)
	sum := 0
	for n := range processedSeedsCount {
		sum += n
	}
	fmt.Printf("Processed %d seeds\n", sum)

}

func process_seed(mapper Mapper, seed int) int {
	num, err := mapper.Transform("seed", "location", seed)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
