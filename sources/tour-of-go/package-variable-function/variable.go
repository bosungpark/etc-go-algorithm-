package main

import "fmt"

var bool_variable bool

func main() {
	var int_variable int
	fmt.Println(bool_variable, int_variable)

	var a, b string= "a", "b"
	fmt.Println(a, b)
	var c, d = "c", "d" // without type
	fmt.Println(c, d)

	// 짧은 변수 선언
	e := "e"
	fmt.Println(e)
}