package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	numbers := []int{4, 6, 9, 45, 8, 17, 3}
	learnerResult, learnerError := CalculateVariance(numbers)
	if learnerError != nil {
		fmt.Println(learnerError)
		return
	}
	fmt.Println(*learnerResult)
}

// CalculateVariance returns variance of the numbers slice, or an error.
func CalculateVariance(numbers []int) (*float64, error) {
	if numbers == nil || len(numbers) == 0 {
		return nil, errors.New("empty slice")
	}

	n := float64(len(numbers))

	sum := func(f func(int) float64) float64 {
		var s float64
		results := make(chan float64, len(numbers))
		for _, num := range numbers {
			go func(i int) {
				results <- f(i)
			}(num)
		}

		for range numbers {
			s += <-results
		}
		return s
	}

	mean := sum(func(i int) float64 {
		return float64(i)
	}) / n

	sumRes := sum(func(i int) float64 {
		return math.Pow(float64(i)-mean, 2)
	})

	variance := sumRes / n

	return &variance, nil
}
