package main

import "fmt"

func main() {
	var s1 = student{name: "wick", grade: 2}

	var s2 *student = &s1
	fmt.Println("student 1, name  :", s1.name)
	fmt.Println("student 4, name  :", s2.name)

	s2.name = "ethan"
	fmt.Println("student 1, name  :", s2.name)
	fmt.Println("student 4, name  :", s2.name)

}

type student struct {
	name  string
	grade int
}
