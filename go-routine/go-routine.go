package main

import (
	"fmt"
	"time"
)

func main() {

	nums := []int{23, 8, 4, 25, 42, 16}
	fmt.Println(sleepSort(nums))

}

func sleepSort(nums[] int) []int {
	ch := make(chan int)
	for _, n := range nums {
		go func(i int) {
			time.Sleep(time.Duration(i) * time.Millisecond)
			ch <- i
		}(n)
	}
	sorted := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sorted[i] = <- ch
	}
	return sorted
}