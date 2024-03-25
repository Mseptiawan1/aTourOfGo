package main

import "fmt"

func main() {
	// var s1 = student{}
	// s1.name = "wawan"
	// s1.grade = 2

	var s2 = student{"ethan", 2}

	// var s3 = student{name: "sasa"}
	// var s4 = student{name: "sasa", grade: 3}

	var s5 = student{grade: 20, name: "sasa"}

	fmt.Println(s2.name)
	fmt.Println(s5.name)

}

type student struct {
	name  string
	grade int
}
