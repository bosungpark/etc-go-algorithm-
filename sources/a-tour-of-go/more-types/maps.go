package main

import "fmt"

// make 함수는 주어진 타입의 초기화되고 사용 준비가 된 맵을 반환합니다.

type Vertex struct {
	Lat, Long float64
}

var m map[string] Vertex

func main() {
	m = make(map[string] Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// 맵 리터럴
	var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
	}
	fmt.Println(m["Google"])

	// same code

	// var m = map[string]Vertex{
	// 	"Bell Labs": {40.68433, -74.39967},
	// 	"Google":    {37.42202, -122.08408},
	// }

	// mutating maps
	m2 := make(map[string]int)

	m2["Answer"] = 42  // set
	fmt.Println("The value:", m2["Answer"])

	m2["Answer"] = 48  // update
	fmt.Println("The value:", m2["Answer"])

	delete(m2, "Answer")  // delete
	fmt.Println("The value:", m2["Answer"])

	v, ok := m2["Answer"]  // check key
	fmt.Println("The value:", v, "Present?", ok)
}