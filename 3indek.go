package main

import "fmt"

func main() {
	var buah = []string{"apple", "grape", "bananase"}
	var buah2 = buah[0:2]
	var buah3 = buah[0:2:2]

	fmt.Println(buah)
	fmt.Println(len(buah))
	fmt.Println(cap(buah))

	fmt.Println(buah2)
	fmt.Println(len(buah2))
	fmt.Println(cap(buah2))

	fmt.Println(buah3)
	fmt.Println(len(buah3))
	fmt.Println(cap(buah3))

}
