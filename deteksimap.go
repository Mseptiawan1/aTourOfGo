package main

import "fmt"

func main() {
	var chicken = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken["januari"]

	if isExist {
		fmt.Println("ada")
		fmt.Println(value)

	} else {
		fmt.Println("item is not exists")
	}

}
