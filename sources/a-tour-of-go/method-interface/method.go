package main

import (
	"fmt"
	"math"
)


type Vertex struct {
	X, Y float64
}

// 메서드는 특별한 receiver 인자가 있는 함수입니다.
// receiver는 func 키워드와 메서드 이름 사이의 자체 인수 목록에 나타납니다.

func (v Vertex) Add() float64 {
	return v.X + v.Y
}

// 구조체가 아닌 형식에 대해서도 메소드를 선언할 수 있습니다.
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Add())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
