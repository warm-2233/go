package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main..start...")

	go func() {
		fmt.Println("func2...")
	}()
	go func() {
		fmt.Println("func3...")
	}()
	go func() {
		fmt.Println("func4...")
	}()
	go func() {
		fmt.Println("func5...")
	}()

	go f1()
	go f2()
	time.Sleep(1 * time.Second)
	fmt.Println("main...end...")
}

func f1() {
	fmt.Println("f1==")
	go func() {
		fmt.Println("func1...")

	}()
}

func f2() {
	go f1()
	fmt.Println("f2==")
}
