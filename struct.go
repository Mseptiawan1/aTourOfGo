package main

import "fmt"

func main() {
	var s1 = student{}
	s1.name = "septiawan"
	s1.grade = 5

	fmt.Sprintln("name : ", s1.name)
	fmt.Sprintln("grade : ", s1.grade)

}

type student struct {
	name  string
	grade int
}
