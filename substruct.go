package main

import "fmt"

func main() {
	var p1 = person{name: "wick", age: 20}
	var p2 = student{person: p1, age: 20, grade: 200}

	fmt.Println(s1.name)
	fmt.Println(s1.age)
	fmt.Println(s1.person.age)

}

type person struct {
	name string
	age  int
}

type student struct {
	person
	age   int
	grade int
}
