package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{43, 5, 3, 1, 3, 4, 5, 6}

	var min, max = getMinMax(numbers)
	fmt.Printf("Data : %v\nmin : %v\nmax : %v\n", numbers, min, max)
}
