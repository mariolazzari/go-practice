package main

import (
	"errors"
	"fmt"
)

func main() {
	numbers := []uint{4, 6, 9, 45, 8, 17, 3}
	learnerResult, learnerError := CalculateVariance(numbers)
	if learnerError != nil {
		fmt.Println(learnerError)
		return
	}
	fmt.Println(*learnerResult)
}

type MeanInput interface {
	int | uint | float64
}

// CalculateVariance returns variance of the numbers slice, or an error.
func CalculateVariance[T MeanInput](numbers []T) (*float64, error) {
	if numbers == nil || len(numbers) == 0 {
		return nil, errors.New("empty slice")
	}

	var sum T
	for _, num := range numbers {
		sum += num
	}

	mean := float64(sum) / float64(len(numbers))

	return &mean, nil
}
