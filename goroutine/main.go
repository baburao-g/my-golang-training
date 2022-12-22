package main

import (
	"fmt"
	"time"
)

// func printNumbers(num int) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(num, ":", i)
// 	}
// }

// func main() {
// 	defer fmt.Println("defer")
// 	fmt.Println(":")
// 	go printNumbers(0)
// 	go printNumbers(1)
// 	time.Sleep(500 * time.Millisecond)
// }

func x(y string) {
	for i := 0; i < 3; i++ {
		fmt.Println(y, ":", i)
	}
}

func main() {
	x("direct")
	go x("routine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")
}
