package main

import "fmt"

func main() {
	// defer문은 자신을 둘러싼 함수가 종료할 때까지 어떠한 함수의 실행을 연기합니다
	// 연기된 호출의 인자는 즉시 평가되지만 그 함수 호출은 자신을 둘러싼 함수가 종료할 때까지 수행되지 않습니다.
	defer fmt.Println("world")

	fmt.Println("hello")

	// 연기된 함수 호출들은 '스택'에 쌓입니다. 한 함수가 종료될 때 그것의 연기된 함수들은 후입선출 순서로 수행됩니다.
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}