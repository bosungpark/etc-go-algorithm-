package main

import "fmt"

func main() {
	i, j := 42, 27

	p:=&i  // point to i
	fmt.Println(*p) // read i through the pointer
	*p=21  // change i by pointer
	fmt.Println(i)  

	
	p = &j
	*p = *p / 3  // change j by pointer
	fmt.Println(j)
}