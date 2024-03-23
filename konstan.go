package main

import "fmt"

func main() {
	const firstname string = "septiawan"

	const fullname = "john wick"
	fmt.Printf("%s versus %s!\n", firstname, fullname)
	fmt.Println("john wick")
	fmt.Println("john", "wick")
	fmt.Print("john wick\n")
	fmt.Print("john ", "wick")
	fmt.Print("john", "wick\n")
}
