package main

import "fmt"

// 4613723
func main() {
	first, second, limit, sum := 1, 2, 4000000, 0
	for value := first + second; value <= limit; {
		if value % 2 == 0 {
			sum += value
		}
		first, second = second, value
	}
	fmt.Println("Sum: ", sum)
}
