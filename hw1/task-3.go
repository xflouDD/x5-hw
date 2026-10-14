package main

import "fmt"

func main() {
	var amount int
	_, err := fmt.Scan(&amount)
	if amount < 0 {
		fmt.Println("Сумма не может быть отрицательной")
		return
	}
	switch {
	case amount < 10000:
		fmt.Println("small")
	case amount < 50000:
		fmt.Println("regular")
	default:
		fmt.Println("large")
	}
}
