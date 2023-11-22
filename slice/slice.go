package main

import (
	"fmt"
	"strings"
)

func main() {
	fruits := []string{"lemon", "orange", "banana"}
	dairy := []string{"cheese", "milk"}
	fmt.Println(concat(fruits, dairy))
	var s1, s2 = "1059", "12"
	fmt.Println(strings.Compare(s1,s2))
}

func concat(s1, s2[]string) []string {
	s3 := make([]string, len(s1) + len(s2))
	copy(s3, s1)
	copy(s3[len(s1):], s2)
	return s3
}
