package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// fmt.Println("Hello World")
	// a := []int{1, 2, 3, 4, -5}
	// for i, v := range a {
	// 	fmt.Println("Index: ", i, "value: ", v)
	// }

	// i := 0
	// for {
	// 	if i >= len(a) {
	// 		break
	// 	}
	// 	fmt.Println("Index:", i, "value:", a[i])
	// 	i++
	// }

	// user input

	// fmt.Printf("Please enter your name: ")
	// var name string
	// fmt.Scanln(&name)
	// fmt.Println("Your name is:", name)

	// command line arguments

	arguments := os.Args
	fmt.Println(arguments)

	var min, max float64
	initialized := 0

	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)

		if err != nil {
			continue
		}

		if initialized == 0 {
			min = n
			max = n
			initialized = 1
			continue
		}
		if n > min {
			max = n
		} else {
			min = n
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
