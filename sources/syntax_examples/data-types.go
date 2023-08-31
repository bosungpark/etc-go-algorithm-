package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var (
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))  //  명시적인 변환을 필요
	// := 혹은 var = 표현을 이용해 명시적인 type을 정의하지 않고 변수를 선언할 때, 그 변수 type은 오른 편에 있는 값으로부터 유추됩니다.
	var z uint = uint(f)  // u := uint(f)
	fmt.Println(x, f, z)
	fmt.Printf("%T, %T, %T, \n", x, f, z)  // int, float64, uint

	// 상수는 := 를 통해 선언될 수 없습니다.
	const Pi = 3.14
	fmt.Printf("%v\n", Pi)  // int, float64, uint

	const (
		// Create a huge number by shifting a 1 bit left 100 places.
		// In other words, the binary number that is 1 followed by 100 zeroes.
		Big = 1 << 100
		// Shift it right again 99 places, so we end up with 1<<1, or 2.
		Small = Big >> 99
	)  // -> 타입이 명시되지 않음
	// type이 정해지지 않은 상수는 그것의 문맥에서 필요한 type을 취합니다.
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

// basic types in go

// bool
// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// int 와 uintptr type은 보통 32-bit 시스템에서는 32 bit, 64-bit 시스템에서는 64 bit의 길이입니다. 
// 정수 값이 필요할 때에는 특정한 이유로 크기를 정해야하거나 unsigned 정수 type을 사용해야하는 게 아니라면 int 를 사용해야합니다.

// byte // uint8의 별칭
// rune // int32의 별칭
// 유니코드에서 code point를 의미합니다.

// float32 float64
// complex64 complex128

// 명시적인 초깃값 없이 선언된 변수는 그것의 zero value 가 주어집니다.

// zero value 는 다음과 같습니다.

// 숫자 type에는 0
// boolean type에는 false
// string에는 "" (빈 문자열)
