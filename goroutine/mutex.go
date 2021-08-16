package goroutine

import (
	"fmt"
	"sync"
)

var id int

func genetateId(mutex *sync.Mutex) int {
	mutex.Lock()
	defer mutex.Unlock()
	id++
	return id
}

func Mutex() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", genetateId(&mutex))
			wg.Done()
		}()
	}
	wg.Wait()
}
