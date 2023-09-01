package main

import (
	"fmt"
	"strings"
) 


func main() {
	// array
	var a [2] string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slice
	var s []int = primes[0:3]
	fmt.Println(s)
	// Slices are like references to arrays
	// 슬라이스는 어떤 데이터도 저장할 수 없습니다. 이것은 단지 기본 배열의 한 영역을 나타낼 뿐입니다.
	// 슬라이스의 요소를 변경하면 기본 배열의 해당 요소가 수정됩니다.
	// 동일한 기본 배열을 공유하는 다른 슬라이스는 이러한 변경사항을 볼 수 있습니다.
	s[0] = 1
	fmt.Println(primes)  // [1 3 5 7 11 13]

	// 슬라이스 리터럴은 길이가 없는 배열 리터럴와 같습니다.
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)  // [2 3 5 7 11 13]

	// 위 배열에서 아래 슬라이스 표현식들은 전부 동일한 의미입니다.
	fmt.Println(q[0:6])  // [2 3 5 7 11 13]
	fmt.Println(q[:6])
	fmt.Println(q[0:])
	fmt.Println(q[:])

	printSlice(q)  // len=6 cap=6 [2 3 5 7 11 13]
	q = q[:0]
	printSlice(q)  // len=0 cap=6 []
	q = q[:4]
	printSlice(q)  // len=4 cap=6 [2 3 5 7]
	// panic: runtime error: slice bounds out of range [:100] with capacity 6
	// q = q[:100]
	// printSlice(q)
	
	// nil slices
	var n [] int
	fmt.Println(n, len(n), cap(n))
	if n == nil {
		fmt.Println("nil!")
	}

	// 슬라이스는 내장된 make 함수로 생성할 수 있습니다. 이것은 동적 크기의 배열을 생성하는 방법입니다.
	m := make([]int, 5)
	printSlice(m)  // len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	printSlice(b)  // len=0 cap=5 []
	
	// 슬라이스는 다른 슬라이스를 포함하여 모든 타입을 담을 수 있습니다. (이차원 배열)
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// 만약 s 의 원래 배열이 너무 작아서 주어진 값을 모두 추가할 수 없을 경우, 더 큰 배열이 할당됩니다. 
	// 이 때 반환된 슬라이스는 새로 할당된 배열을 가리킵니다.
	b = append(b, 1)
	printSlice(b)  // len=1 cap=5 [1]
	b = append(b, 2, 3, 4, 5, 6)
	printSlice(b)  // len=6 cap=10 [1 2 3 4 5 6]

}

func printSlice(s []int) {
	// cap: 기본 배열의 요소의 개, len: 슬라이스가 포함하는 요소의 개수
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}