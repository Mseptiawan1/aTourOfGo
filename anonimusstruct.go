package main

import "fmt"

func main() {
	var s1 = struct {
		person
		grade int
	}{}

	s1.person = person{"wick", 20}
	s1.grade = 200

	fmt.Println("name  : ", s1.person.name)
	fmt.Println("age  : ", s1.person.age)
	fmt.Println("grade : ", s1.grade)

}

type person struct {
	name string
	age  int
}
