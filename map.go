package main

import "fmt"

func main() {

	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *map[string]int)

	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) //50
	fmt.Println("mei", chicken["mei"])         //0

	fmt.Println("januari", chicken1["januari"]) //50
	fmt.Println("mei", chicken1["mei"])         //0

	fmt.Println("januari", chicken2["januari"]) //50
	fmt.Println("mei", chicken2["mei"])         //0

}
