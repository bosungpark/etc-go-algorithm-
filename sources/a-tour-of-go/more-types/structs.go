package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})  // X -> 1, Y -> 2

	v := Vertex{1, 2}
	fmt.Println(v.X)  // 1

	p := &v
	p.X = 1e9  // (*p).X = 1e9
	fmt.Println(v.X)  // 1000000000

	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p1  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, v2, v3, p1)
}