package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	var m sync.Mutex

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			fmt.Println(incrementer)
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value", incrementer)
}
