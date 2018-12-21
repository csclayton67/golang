package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Num CPU's", runtime.NumCPU())
	fmt.Println("Num Go", runtime.NumGoroutine())
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("Hello Thing One")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello Thing Two!")
		wg.Done()
	}()
	fmt.Println("Mid CPU's", runtime.NumCPU())
	fmt.Println("Mid Go", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("End CPU's", runtime.NumCPU())
	fmt.Println("End Go", runtime.NumGoroutine())
	fmt.Println("About to exit!")
}
