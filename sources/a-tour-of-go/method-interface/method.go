package main

// 메서드는 특별한 receiver 인자가 있는 함수입니다.
// receiver는 func 키워드와 메서드 이름 사이의 자체 인수 목록에 나타납니다.

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Add() float64 {
	return v.X + v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Add())
}
