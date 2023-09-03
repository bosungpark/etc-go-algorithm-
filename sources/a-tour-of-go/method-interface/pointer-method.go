package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Add() float64 {
	return v.X + v.Y
}
// pointer를 사용하는 것은 스트럭쳐의 인자를 수정할 수 있게 해준다.
// pointer가 없다면 복사본을 수정하는 것이라, 실제 스트럭쳐의 값은 변하지 않는다.
// 사용자의 실수가 있더라도 value receiver기준으로 통일된다. (indirection)
func (v *Vertex) Scale(f float64) {  
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Add()) // 70 (pointer를 사용하지 않으면 7)
}