package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func RunOne() {
	file, err := os.Open(FILE)
	if err != nil {
		log.Fatal("Error opening file")
	}
	defer file.Close()
	fmt.Printf("Worker count is %d\n", WorkerCount)
	scanner := bufio.NewScanner(file)
	inputChan := make(chan string, WorkerCount)
	pointsChan := make(chan int, WorkerCount)
	var wg sync.WaitGroup
	done := make(chan bool)
	go func() {
		for scanner.Scan() {
			inputChan <- scanner.Text()
		}
		close(inputChan)
	}()
	for i := 0; i < WorkerCount; i++ {
		wg.Add(1)
		go processInputs(&wg, inputChan, pointsChan)
	}
	go func() {
		sum := 0
		for points := range pointsChan {
			sum += points
		}
		fmt.Printf("Sum: %d\n", sum)
		done <- true
	}()
	wg.Wait()
	close(pointsChan)
	<-done
}

func processInputs(wg *sync.WaitGroup, inputChan <-chan string, pointsChan chan<- int) {
	defer wg.Done()
	for line := range inputChan {
		card := NewCard(line)
		pointsChan <- card.CalculatePoints()
	}
}
