package main

import (
	"sync"
	"testing"
)

func TestAdd(t *testing.T) {
	numberGoroutines := 10

	var wg sync.WaitGroup
	var counter int
	var mu sync.Mutex

	randomSlice := GenerateRandomSlice(10000000, 1, 100)

	sum := 0

	for _, value := range randomSlice {
		sum += value
	}

	wg.Add(numberGoroutines)

	for i := 0; i < numberGoroutines; i++ {
		go Add(randomSlice, &counter, &mu, &wg, i, numberGoroutines)
	}

	wg.Wait()

	if sum != counter {
		t.Errorf("counter returned %d, expected %d", counter, sum)
	}
}