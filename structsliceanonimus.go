package main

import "fmt"

func main() {
	var allStudents = []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 21},
		{person: person{"ethan", 100}, grade: 33},
		{person: person{"lisa", 50}, grade: 54},
	}

	for _, stsudent := range allStudents {
		fmt.Println(student)
	}

}

type person struct {
	name string
	age  int
}
