package main

import "fmt"

func main() {
	var fruits = make([]string, 2)
	fruits[0] = "apple"
	fruits[1] = "manggo"

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])

	}
}
