package main

import "fmt"

func main() {
	size := 1

	Loop:
		for {
			switch {
			case size == 10:
				fmt.Println("ending")
				break Loop
			default:
				size += 1
			}
		}
}
