package main

import (
	"fmt"
	"sync"
)

func pow(num int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	result <- num * num
}

func main() {
	var wg sync.WaitGroup
	slice := []int{2, 4, 6, 8}
	result := make(chan int)
	for pos := range slice {
		wg.Add(1)
		pow(slice[pos], result, &wg)
	}
	go func(result chan int, wg *sync.WaitGroup) {
		wg.Wait()
		close(result)
	}(result, &wg)
	sum := 0
	for res := range result {
		sum += res
	}
	fmt.Println(sum)
}
