package main

import "fmt"

func main()()  {
	v1 := []int{1, 2, 3}
	v2 := []int{1, 2, 3}
	fmt.Println("v1.v2: ", dot(v1, v2))
}

// v1.v2 = v1[0]*v2[0] + v1[1]*v2[1]....
func dot(v1 []int, v2 []int) int {
	sum := 0
	for i, n := range v1 {
		sum += n * v2[i]
	}
	return sum
}