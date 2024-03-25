package main

import "fmt"

func main() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d :  %s\n ", i+1, fruits[i])
	}

	var buah = [4]string{"pisang", "jeruk", "mangga", "apel"}

	/* for i, buah := range buah {
		fmt.Printf("elemen %d :  %s\n ", i, buah)
	} */

	for _, buah := range buah {
		fmt.Printf("elemen buah :  %s\n ", buah)
	}


	for i,_ := rang buah{
		// blablabla
	}
}
