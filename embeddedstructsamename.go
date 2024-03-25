package main

import "fmt"

func main() {
	var s1 = student{}

	s1.age = 21        //struct student
	s1.person.age = 22 // akses struct person

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
