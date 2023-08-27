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

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))  // 55
	fmt.Println(minus(42, 12))  // 30

	a, b := swap("hello", "world") // fmt.Println(swap("hello", "world"))
	fmt.Println(a, b)

	fmt.Println(split(17)) // 7 10
}