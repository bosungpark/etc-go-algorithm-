package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func minus(x, y int) int {
	return x - y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(42, 13))  // 55
	fmt.Println(minus(42, 12))  // 30

	a, b := swap("hello", "world") 
	fmt.Println(a, b)
}