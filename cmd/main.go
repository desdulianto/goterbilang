package main

import (
	"fmt"

	terbilang "github.com/desdulianto/goterbilang"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Println(terbilang.FromInt(i))
	}
	for i := 101; i <= 1000; i += 3 {
		fmt.Println(terbilang.FromInt(i))
	}
	for i := 1000; i <= 100000; i += 3333 {
		fmt.Println(terbilang.FromInt(i))
	}
}
