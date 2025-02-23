package main

import (
	"errors"
	"math"
)

func main() {
	numbers := []int{4, 6, 9, 45, 8, 17, 3}
	learnerResult, learnerError := CalculateVariance(numbers)
	if learnerError != nil {
		panic(learnerError)
	}
	println(*learnerResult)
}

// CalculateVariance returns variance of the numbers slice, or an error.
func CalculateVariance(numbers []int) (*float64, error) {
	if numbers == nil || len(numbers) == 0 {
		return nil, errors.New("empty slice")
	}

	n := float64(len(numbers))
	sum := func(f func(int) float64) (s float64) {
		for _, num := range numbers {
			s += float64(f(num))
		}
		return s
	}

	mean := sum(func(num int) float64 {
		return float64(num)
	}) / n

	sumRes := sum(func(i int) float64 {
		return math.Pow(float64(i)-mean, 2)
	})

	variance := sumRes / n

	return &variance, nil
}
