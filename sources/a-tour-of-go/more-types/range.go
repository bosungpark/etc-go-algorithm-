package main

import "fmt"

// 슬라이스에서 range 를 사용하면, 각 순회마다 두 개의 값이 반환됩니다. 
// 첫 번째는 인덱스이고, 두 번째는 해당 인덱스 값의 복사본입니다.

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	fmt.Printf("\n")
	// 만약 인덱스만을 원하면, 두 번째 변수를 생략할 수 있습니다.
	for i := range pow {
		fmt.Printf("%d\n", i)
	}
	fmt.Printf("\n")
	// _ 을 할당하여 인덱스 또는 값을 건너뛸 수 있습니다.
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
