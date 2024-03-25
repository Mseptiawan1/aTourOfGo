package main

import "fmt"

func main() {
	var fruits = []string{"apple", "Grape", "banana", "melon"}
	fmt.Println(fruits[0])

	// the connection slice and array
	var fruits2 = []string{"apple", "Grape", "banana", "melon"}
	var newFruits = fruits2[0:2]

	fmt.Println(newFruits)
}
