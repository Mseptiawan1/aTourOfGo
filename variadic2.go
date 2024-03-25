package main

import (
	"fmt"
	"strings"
)

func main() {
	yourHobbies("septiawan", "mancing", "gitaran", "menyusu")
	var myHobbies = []string{"bernyanyi", "bersiul"}
	yourHobbies("septiawan", myHobbies...)

}
func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("hello, my name is : %s\n", name)
	fmt.Printf("my hobbies are : %s\n", hobbiesAsString)

}
