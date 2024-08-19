package main

import (
	"fmt"
	"sync"
	"time"
)

func printOne(wg *sync.WaitGroup){
	defer wg.Done()

	for i := 0; i < 50; i++ {
		fmt.Print("1")
		time.Sleep(10 * time.Millisecond)
	}
}

func printTwo(wg *sync.WaitGroup){
	defer wg.Done()

	for i := 0; i < 50; i++ {
		fmt.Print("2")
		time.Sleep(10 * time.Millisecond)
	}
}

func sum(numbers []int, resultChan chan <- int, wg *sync.WaitGroup){
	defer wg.Done()

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	resultChan <- sum
}

func main() {
	var wg sync.WaitGroup

	resultChan := make(chan int) //<-
	numbers := []int{1,2,3,4,5,6,7,8,9,10}

	numGoRoutines := 4

	sliceSize := len(numbers) / numGoRoutines

	for i := 0; i < numGoRoutines; i++ {
		startIdx := i * sliceSize
		endIdx := (i + 1) * sliceSize

		if i == numGoRoutines-1 {
			endIdx = len(numbers)
		}

		wg.Add(1)
		go sum(numbers[startIdx:endIdx], resultChan, &wg)
	}

	var collectWg sync.WaitGroup
	collectWg.Add(1)

	go func(){
		defer collectWg.Done()
		totalSum := 0

		for i := 0; i < numGoRoutines; i++ {
			partialSum := <-resultChan //receiving the value
			totalSum += partialSum
		}
		fmt.Printf("Sum of numbers %d\n", totalSum)
	}()


	// wg.Add(2)
	// go printOne(&wg)
	// go printTwo(&wg)

	wg.Wait()
	close(resultChan)
	collectWg.Wait()
}
