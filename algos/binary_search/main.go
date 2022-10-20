package main

import (
	"fmt"
	"math"
	"math/rand"
)

// estimating the mean via binary search
func binarySearchMean(v []float64) (float64, int) {
	low := minimum(v)
	high := maximum(v)
	mn := mean(v)
	sens := .001
	guess := (low + high) / 2
	counter := 1

	res := math.Abs(guess - mn)

	for res >= sens {
		fmt.Printf("Counter = %v; Guess = %v\n", counter, guess)
		if guess < mn {
			low = guess
		} else {
			high = guess
		}

		counter++
		guess = (low + high) / 2
		res = math.Abs(guess - mn)
	}

	return guess, counter

}

func maximum(v []float64) float64 {

	var m float64

	for i, e := range v {
		if i == 0 || e > m {
			m = e
		}
	}

	return m
}

func minimum(v []float64) float64 {
	var m float64
	for i, e := range v {
		if i == 0 || e < m {
			m = e
		}
	}

	return m
}

func mean(v []float64) float64 {
	l := len(v)
	var sum float64

	if l == 0 {
		return 0.0
	}

	for _, d := range v {
		sum += d
	}

	m := sum / float64(l)

	return m
}

func main() {
	var v []float64

	s := 100

	for i := 1; i <= s; i++ {
		j := rand.NormFloat64()

		v = append(v, j)
	}

	m := maximum(v)

	fmt.Println(m)

	r, c := binarySearchMean(v)

	mn := mean(v)

	fmt.Printf("Guessed %v after %v iterations\n", r, c)
	fmt.Printf("Mean is actually %v", mn)

}
