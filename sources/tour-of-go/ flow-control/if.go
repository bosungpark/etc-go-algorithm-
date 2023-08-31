package main

import "fmt"

func isOdd(x int) bool {
	if x % 2 == 1 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isOdd(1), isOdd(2))

	if v:=isOdd(2); v == true {
		fmt.Println("simple if statment:", v)
	} else {
		fmt.Println("simple if statment:", v)
	}
}