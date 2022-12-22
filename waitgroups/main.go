package main

// import (
// 	"fmt"
// 	"sync"
// )

// var x = 0

// func routine(i int, wg *sync.WaitGroup) {
// 	fmt.Println("started routine", i)
// 	// time.Sleep(time.Second * 2)
// 	fmt.Printf("Routine %d routine\n", i)
// 	wg.Done()
// }

// func main() {
// 	noOfRoutines := 3
// 	var wg sync.WaitGroup

// 	for i := 0; i < noOfRoutines; i++ {
// 		wg.Add(1)
// 		go routine(i, &wg)
// 	}
// 	wg.Wait()
// 	fmt.Println("All routines are finished")
// }

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for j := range jobs {
// 		// fmt.Println("worker", id, "started job", j)
// 		// time.Sleep(time.Second)
// 		// fmt.Println("worker", id, "finished job", j)
// 		results <- j * 2
// 	}
// }

// func main() {
// 	start := time.Now()
// 	const numJobs = 10000

// 	jobs := make(chan int, numJobs)
// 	results := make(chan int, numJobs)
// 	for w := 1; w <= 12; w++ {
// 		go worker(w, jobs, results)
// 	}

// 	for j := 1; j <= numJobs; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)

// 	for a := 1; a <= numJobs; a++ {
// 		<-results
// 	}

// 	end := time.Since(start)
// 	fmt.Println("time:", end)
// }

// --------------------
// type singleTon struct {
// 	connectionString string
// }

// var lock = &sync.Mutex{}
// var instance *singleTon

// func getInstance() *singleTon {
// 	lock.Lock()
// 	defer lock.Unlock()
// 	if instance == nil {
// 		fmt.Println("hi1")

// 		instance = &singleTon{connectionString: "dbConnection"}
// 	}
// 	return instance
// }

// func main() {
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			time.Sleep(time.Millisecond * 30)
// 			fmt.Println(getInstance(), "-", i)
// 		}
// 	}()

// 	go func() {
// 		fmt.Println("hi")
// 		inst := getInstance()
// 		fmt.Println(inst.connectionString)

// 	}()

// 	time.Sleep(time.Millisecond * 1000)

// }

// --------------------
