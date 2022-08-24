package main

import (
	"fmt"
	"sync"
)

func main() {

	var (
		wg      sync.WaitGroup
		mu      sync.Mutex
		counter int
	)

	for i := 0; i < 50; i++ {

		wg.Add(1)

		go func() {

			defer wg.Done()

			for j := 0; j < 1000; j++ {

				// Sem o mutex lockando a variável counter, as goroutines irão alterar o conteúdo
				// dela ao mesmo tempo, resultando assim em valores inconsistenctes no final da função.
				// O mutex garante que apenas um goroutine por vez irá alterar o conteúdo da variável.

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
