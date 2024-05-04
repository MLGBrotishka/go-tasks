package main

import (
	"fmt"
	"sync"
)

func pow(num *int, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := (*num)
	(*num) = temp * temp
}

func main() {
	var wg sync.WaitGroup
	slice := []int{2, 4, 6, 8}
	for pos := range slice {
		wg.Add(1)
		pow(&slice[pos], &wg)
	}
	wg.Wait()
	for pos := range slice {
		fmt.Println(slice[pos])
	}
}
