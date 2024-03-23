package main

import "fmt"

func main() {
	var point = 6
	switch point {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	var titik = 10
	switch titik {
	case 8:
		fmt.Println("perfect")
	case 9, 10, 3, 44:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
	var titik2 = 10
	switch titik2 {
	case 8:
		fmt.Println("perfect")
	case 9, 10, 3, 44:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("not bad kali")

		}
	}

}
