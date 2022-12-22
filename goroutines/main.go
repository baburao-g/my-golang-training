package main

import (
	"fmt"
	"time"
)

// func main() {
// 	message := make(chan string)

// 	go func() {

// 		message <- "ping"
// 		fmt.Println("done")
// 	}()

// 	time.Sleep(5 * time.Second)
// 	// msg := <-message
// 	// fmt.Println(msg)

// 	// msg = <-message
// 	fmt.Println("msg")
// }

func myFunc(mychan chan int) {
	for i := 0; i < 4; i++ {
		time.Sleep(time.Second)
		mychan <- i
	}
	close(mychan)
}

func main() {
	c := make(chan int)
	go myFunc(c)

	// for {
	// 	res, ok := <-c
	// 	if !ok {
	// 		fmt.Println("channel closed:", ok)
	// 		break
	// 	}
	// 	fmt.Println("channel open:", res, ok)

	// }
	fmt.Println("Start")
	for res := range c {
		fmt.Println(res)
	}
	fmt.Println("Done")
}

// func write(ch chan int) {
// 	for i := 0; i < 100; i++ {
// 		ch <- i
// 		time.Sleep(time.Second)
// 		fmt.Println("Written to channel:", i)
// 	}
// 	close(ch)
// // }

// func main() {
// 	c := make(chan int)
// 	c <- 1
// 	// fmt.Println(<-c)
// 	// fmt.Println(<-c)
// 	fmt.Println(<-c)

// go write(c)
// fmt.Println("Start")
// for res := range c {
// 	fmt.Println("Read value from channel:", res)
// 	time.Sleep(time.Second)
// }
// fmt.Println("Done")
// }

// func main() {
// 	marks := [5]int{1, 2, 3, 4, 5}
// 	ch := make(chan int)
// 	go func() {
// 		var sum int
// 		for res := range marks {
// 			sum += res
// 		}
// 		ch <- sum
// 		close(ch)
// 	}()
// 	fmt.Println(<-ch)
// }

// func sum(ch chan int, num int)  {
// 	ch
// }
// func main() {
// 	marks := [5]int{1, 2, 3, 4, 5}
// 	ch := make(chan int)

// 		for res := range marks {
// 			sum += sum(ch)
// 	fmt.Println(<-ch)
// }
