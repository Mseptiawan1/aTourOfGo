package main

import "fmt"

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int

	for _, e := range numbers {
		if e <= max {
			res = append(res, e)

		}
	}
	return len(res), func() []int {
		return res
	}
}

func main() {
	var max = 3
	var numbers = []int{3, 5, 6, 4, 2, 1, 4, 5, 6, 7}
	var howMany, getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t :", numbers)
	fmt.Println("find\t : %d\n\n", max)

	fmt.Println("found\t : ", howMany)
	fmt.Println("value\t : %d\n\n", theNumbers)

}
