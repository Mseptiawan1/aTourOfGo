package main

import "fmt"

func main() {

	var names [4]string
	names[0] = "m"
	names[1] = "t"
	names[2] = "b"
	names[3] = "a"

	fmt.Println(names[0], names[1], names[2], names[3])

	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("jumlah elemen \t\t", len(fruits))
	fmt.Println("isi semua elemen \t", fruits)

	// var fruits2 = [4]string{
	// 	"apple",
	// 	"grape",
	// 	"banana",
	// 	"melon",
	// }

	var bilangan = [...]int{1, 2, 3, 4, 5}
	fmt.Println("data array \t :", bilangan)
	fmt.Println("jumlah elemen  \t :", len(bilangan))

	var bilangan1 = [2][3]int{{1, 2, 3}, {66, 44, 22}}
	fmt.Println("bilangan1 : ", bilangan1)
	var bilangan2 = [2][3]int{[3]int{3, 3, 2}, [3]int{3, 4, 5}}
	fmt.Println("bilangan2: ", bilangan2)

}
