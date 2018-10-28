package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello world1!")
	}()
	go func() {
		fmt.Println("Hello world2!")
	}()
	time.Sleep(time.Second * 3)
}
