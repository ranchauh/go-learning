package main

import "fmt"

func main() {
	maxlen, result := 0, 0
	for i := 1; i < 10_00_000; i++ {
		if length := findChain(i); length > maxlen {
			maxlen, result = length, i
		}
	}
	fmt.Printf("Starting number %v produced the longest chain of length %v", result, maxlen)
}

func findChain(num int) int  {
	length := 0
	for num != 1 {
		num = collatz(num)
		length++
	}
	return length
}

func collatz(num int)  int {
	if num % 2 == 0 {
		return num/2
	} else {
		return 3*num + 1
	}
}
