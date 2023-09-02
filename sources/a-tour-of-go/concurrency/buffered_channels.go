package main

import "fmt"

func main() {
	// buffer가 없으면 데드락 발생(fatal error: all goroutines are asleep - deadlock!)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
