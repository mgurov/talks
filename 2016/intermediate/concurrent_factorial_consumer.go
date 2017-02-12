package main

import (
	"fmt"
	"sync"
)

// START OMIT

func main() {
	result := make(chan int)
	wg := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go produce(i, result, &wg)
	}
	go consume(result)
	wg.Wait()
	close(result)
}

func produce(i int, c chan int, wg *sync.WaitGroup) {
	c <- fac(i)
	wg.Done()
}

func consume(result chan int) {
	for r := range result {
		fmt.Println(r)
	}
}

// END OMIT

func fac(i int) int {
	if i == 1 {
		return 1
	} else {
		return i * fac(i-1)
	}
}
