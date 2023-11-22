package main

import (
	"errors"
	"fmt"
	"sort"
)

func main1() {
	values := []float64{3.0}
	m, err := median(values)
	if err != nil {
		fmt.Println("Error: ", err)
	}else {
		fmt.Printf("%-0.2f -> %-0.2f\n", values, m)
	}
}

func median(values []float64) (float64, error){
	size := len(values)
	if size == 0 {
		return 0, errors.New("values is empty")
	}
	sort.Float64s(values)
	i := size/2
	if size%2 != 0 {
		return values[i], nil
	} else {
		return (values[i-1] + values[i])/2, nil
	}
}