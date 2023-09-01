package main

import (
	"fmt"
)


func main() {

	sum := func(x, y int) int {
		return x + y
	}

	sub := func(x, y int) int {
		return x - y
	}

	fmt.Println(compute(sum))  // 7
	fmt.Println(compute(sub))  // -1

	// closer
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// 함수 값은 함수의 인수나 반환 값으로 사용될 수 있습니다.
func compute(fn func(int, int) int) int {
	return fn(3, 4)
}

// 클로저는 함수의 외부로부터 오는 변수를 참조하는 함수 값
// 각 클로저는 그 자체의 sum 변수에 bound(바운드) 되어 있습니다.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
