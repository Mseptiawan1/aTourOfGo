package main

import "fmt"

func main() {
	var p2 = struct {
		name string
		age  int
	}{age: 20, name: "wick"}
	var p3 = struct {
		name string
		age  int
	}{"ethan", 24}

	fmt.Println(p2.name)
	fmt.Println(p3.name)

}

type student struct {
	person struct {
		name string
		age  int
	}
	grade   int
	hobbies []string
}

// horizontal struct
type person2 struct {
	name    string
	age     int
	hobbies []string
}

type person3 struct {
	name string `tag1`
	age  string `tag2`
}
