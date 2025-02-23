package main

import (
	"errors"
	"fmt"
)

func main() {
	avg, err := CalculateMean([]int{1, 2, 3})
	printResults(avg, err)

	avg, err = CalculateMean([]int{})
	printResults(avg, err)

	avg, err = CalculateMean([]int{123, 123, 2, 3})
	printResults(avg, err)
}

func printResults(avg *float64, err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Println("mean:", *avg)
}

// CalculateMean returns mean of the numbers slice, or an error.
func CalculateMean(numbers []int) (*float64, error) {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	if sum == 0 {
		return nil, errors.New("empty array")
	}
	mean := float64(sum / len(numbers))

	return &mean, nil
}
