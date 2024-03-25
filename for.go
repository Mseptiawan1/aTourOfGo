package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("angka", i)
	}

	i := 0
	for i < 5 {
		fmt.Println("angka", i)
		i++
	}

	i = 0
	for {
		fmt.Println("angka", i)
		i++
		if i == 5 {
			break

		}
	}
}
