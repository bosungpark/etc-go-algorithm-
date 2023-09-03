package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe_empty_interface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I

	describe(i)  // (<nil>, <nil>)
	// i.M()  // panic: runtime error: invalid memory address or nil pointer dereference

	var t *T
	i = t
	describe(i)  // (<nil>, *main.T)
	i.M()  // <nil>

	i = &T{"hello"}  
	describe(i)  // (&{hello}, *main.T)
	i.M()  // hello

	// 빈 인터페이스는 모든 유형의 값을 가질 수 있습니다. (모든 유형은 최소 0개의 메소드를 구현합니다.)
	var e interface{}
	describe_empty_interface(e)  // (<nil>, <nil>)

	e = 42
	describe_empty_interface(e)  // (42, int)

	e = "hello"
	describe_empty_interface(e)  // (hello, string)
}
