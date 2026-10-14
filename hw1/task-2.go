package main

import "fmt"

func stats(orders []int) (int, int, int, float64) {
	if len(orders) == 0 {
		return 0, 0, 0, 0
	}
	min := orders[0]
	max := orders[0]
	sum := 0
	for _, c := range orders {
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
		sum += amount
	}
	avg := float64(sum) / float64(len(orders))
	return min, max, sum, avg
}
