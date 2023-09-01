package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	// 초기화 구문과 사후 구문은 필수는 아닙니다.
	sum = 1
	for ; sum * 2 < 10; {
		sum += sum
	}
	fmt.Println(sum)

	// ; 을 생략할 수 있다는 점에서 C의 while 은 Go에서 for 로 쓰입니다.
	sum = 1
	for sum * 2 < 10 {
		sum += sum
	}
	fmt.Println(sum)

	// 반복 조건을 생략하면 간단하게 무한 루프를 만들 수 있습니다.
	// for {
	// 	fmt.Println("inf loop")
	// }
}