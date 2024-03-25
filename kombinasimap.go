package main

import "fmt"

func main() {
	var data = []map[string]int{
		{"name": "septiawan", "umur": 20},
		{"alamat": "yuka", "hobi": "bola"},
	}
	fmt.Println(data["name"])
}
