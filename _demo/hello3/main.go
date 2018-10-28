package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("Hello world1!")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println("Hello world2!")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println("Hello world3!")
		wg.Done()
	}()
	wg.Wait()

}
