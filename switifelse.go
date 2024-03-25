package main

import "fmt"

func main() {
	var point = 6
	switch {
	case point == 6:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("lumayan")
	default:
		{
			fmt.Println("noob")
			fmt.Println("bodo")

		}
	}
}
