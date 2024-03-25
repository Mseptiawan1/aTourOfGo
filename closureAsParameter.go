package main

import (
	"fmt"
	"strings"
)

type Filtercallback func(string) bool

func filter(data []string, callback Filtercallback) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)

		}

	}
	return result

}

func main() {
	var data = []string{"wick", "jason", "ethan"}
	var dataContains0 = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var dataLenght5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("Data asli \t\t:", data)

	fmt.Println("filter data huruf \"i\"t:", dataContains0)
	fmt.Println("filter jumlah huruf \"5\"t:", dataLenght5)

}
