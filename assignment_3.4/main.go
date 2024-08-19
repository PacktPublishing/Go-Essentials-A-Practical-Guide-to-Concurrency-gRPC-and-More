// Your task is to create a Go function that calculates the sum of a given slice of integers concurrently using goroutines and ensures thread-safety using a mutex. Additionally, you will write a test function to verify the correctness and safety of the concurrent sum calculation. Try to account for situations where the amount of goroutines to values in the slice is not perfectly divisible. Example: 8 goroutines created for 100 values in the slice.

package main

import (
	"math/rand"
	"sync"
	"time"
)

func GenerateRandomSlice(size, min, max int) []int{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomSlice := make([]int, size)
	for i := 0; i < size; i++ {
		randomSlice[i] = r.Intn(max-min+1) + min
	}

	return randomSlice
}

func Add(numbers []int, counter *int, mu *sync.Mutex, wg *sync.WaitGroup, goroutineIndex, numberGoroutines int){
	defer wg.Done()

	numEl := len(numbers)

	elementsPerGoroutine := (numEl + numberGoroutines - 1) / numberGoroutines

	start := goroutineIndex * elementsPerGoroutine
	end := (goroutineIndex + 1) * elementsPerGoroutine

	if end > numEl {
		end = numEl
	}

	mu.Lock()

	for i:=start; i < end; i++{
		*counter += numbers[i]
	}
	mu.Unlock()
}

func main() {}