package main

import (
	"fmt"
	"sync"
)

func increment(counter *int, mu *sync.Mutex){
	mu.Lock()
	defer mu.Unlock()
	*counter++
}

func main() {
	numGoRoutines := 1000
	var wg sync.WaitGroup
	var counter int
	var mu sync.Mutex

	for i := 0; i < numGoRoutines; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			increment(&counter, &mu)
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}
