package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("Hello world1!")
		ch<-1
	}()
	go func() {
		fmt.Println("Hello world2!")
		ch<-2
	}()
	go func() {
		fmt.Println("Hello world3!")
		ch<-3
	}()
	go func() {
		fmt.Println("Hello world4!")
		ch<-4
	}()
	for i:=0; i<4; i++ {
		hello := <-ch
		fmt.Println("done", hello)
	}
}
