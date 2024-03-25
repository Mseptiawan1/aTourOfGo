package main

import "fmt"

func main() {
	var buah = []string{"apple", "grape", "banana"}
	var buah2 = append(buah, "pepaya")
	fmt.Println(len(buah))

	fmt.Println(buah)
	fmt.Println(buah2)
	fmt.Println(cap(buah2))

	buahan := []string{"apple", "Grape", "banana"}
	var buahan2 = buahan[0:2]
	fmt.Println(cap(buahan))
	fmt.Println(+len(buahan2))
	fmt.Println("atas")

	fmt.Println(buahan)
	fmt.Println(buahan2)

	var cbuahan = append(buahan2, "pepayya")

	fmt.Println(buahan)
	fmt.Println(buahan2)
	fmt.Println(cbuahan)

}
