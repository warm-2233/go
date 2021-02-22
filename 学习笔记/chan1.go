package main

import (
	"fmt"
	// "time"
)

func main() {
	ch := make(chan int)
	fmt.Println("main start..")

	go func() {
		fmt.Println("func..")
		// a := <-ch
		// fmt.Println(a)
		ch <- 1
		fmt.Println("func end..")
	}()
	// time.Sleep(30 * time.Second)
	fmt.Println(<-ch)
	fmt.Println("main end..")
}
