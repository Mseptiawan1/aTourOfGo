package main

import "fmt"

func main() {
	var s1 = student{}
	s1.name = "wick"
	s1.age = 21
	s1.grade = 2

	fmt.Println("name : ", s1.name)
	fmt.Println("age : ", s1.age)
	fmt.Println("name : ", s1.person.age)
	fmt.Println("name : ", s1.grade)

}

type person struct {
	name string
	age  int
}

type student struct {
	grade int
	person
}
