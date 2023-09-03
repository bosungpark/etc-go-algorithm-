package main

import "fmt"

// 가장 널리 사용되는 인터페이스 중 하나는 fmt 패키지에 의해 정의된 Stringer 입니다.

// type Stringer interface {
//     String() string
// }

// Stringer 는 자신을 문자열로 설명할 수 있는 타입입니다. 
// fmt 패키지 및 기타 여러 패키지는 값을 출력하기 위해 이 인터페이스를 사용합니다.

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
