package main

import "fmt"

func main() {
	var allStudents = []person{
		{name: "wick", age: 21},
		{name: "ethan", age: 24},
		{name: "bourne", age: 22},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

}

type person struct {
	name string
	age  int
}
