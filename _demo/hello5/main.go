package main

import (
	"sync"
	"fmt"
)

func main() {
	wg := sync.WaitGroup{}
	chTom := make(chan string)
	chJerry := make(chan string)

	wg.Add(1)
	// Tom task
	go func() {
		for {
			w := <-chTom
			if w == "Hello Tom" {
				fmt.Println("Tom : Yeah, Hello Jerry.")
				chJerry <- "Hello Jerry"
			} else {
				fmt.Println("Tom : Nice Day.")
				chJerry <- "Nice Day"
				wg.Done()
				break
			}
		}
	}()

	wg.Add(1)
	// Jerry task
	go func() {
		for {
			w := <-chJerry
			if w == "Hello Jerry" {
				fmt.Println("Jerry : Fine, Tom.")
				chTom <- "Fine Tom"
			} else {
				wg.Done()
				break
			}
		}
	}()

	//fmt.Println("Jerry : Hello Tom.")
	//chTom <- "Hello Tom"
	fmt.Println("Tom : Hello Jerry.")
	chJerry <- "Hello Jerry"
	wg.Wait()
}
