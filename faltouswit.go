package main

import "fmt"

func main() {
	point, kedua := 6, 10

	switch {
	case point == 6:
		fmt.Println("perfect")
		fallthrough

	case (kedua < 8) && (kedua > 3):
		fmt.Println("lumayan")
	default:
		{

			fmt.Println("noob")

		}
	}
}
