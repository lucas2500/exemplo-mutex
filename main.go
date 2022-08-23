package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter int

	for i := 0; i < 50; i++ {

		wg.Add(1)

		go func() {

			defer wg.Done()

			for j := 0; j < 1000; j++ {

				mu.Lock()
				// Seção crítica
				counter++
				mu.Unlock()
			}

		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
