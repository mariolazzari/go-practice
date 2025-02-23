package main

import "errors"

func main() {
	avg, err := CalculateMean([]int{1, 2, 3})
	if err != nil {
		println("error", err)
		return
	}
	print("mean:", *avg)
}

// CalculateMean returns mean of the numbers slice, or an error.
func CalculateMean(numbers []int) (*float64, error) {
	// Your code goes here.
	mean := 0.0
	for _, n := range numbers {
		mean += float64(n)
	}
	if mean == 0 {
		return nil, errors.New("empty array")
	}
	mean /= float64(len(numbers))

	return &mean, nil
}
