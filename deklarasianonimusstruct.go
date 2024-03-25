package main

import "fmt"

func main() {

	var student struct {
		person
		grade int
	}
	student.person = person{"wick", 21}
	student.grade = 2

	fmt.Println(student.grade)
	fmt.Println(student.person.name)
	fmt.Println(student.person.age)

}

type person struct {
	name string
	age  int
}
