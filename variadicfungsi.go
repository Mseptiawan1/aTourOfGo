package main

import "fmt"

func main() {
	var avg = calculate(2, 4, 5, 6, 7, 4, 7)
	var msg = fmt.Sprintf("rata-rata : %.2f", avg)
	fmt.Println(msg)

	// slice
	var numbers2 = []int{4, 5, 6, 2, 4, 5, 6, 24, 6}
	var avg2 = calculate(numbers2...)
	var msg2 = fmt.Sprintf("rata-rata : %.2f", avg2)
	fmt.Println(msg2)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers))

	return avg
}
