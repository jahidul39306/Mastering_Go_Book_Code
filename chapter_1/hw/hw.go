package main

import "fmt"

func main() {
	// fmt.Println("Hello World")
	a := []int{1, 2, 3, 4, -5}
	// for i, v := range a {
	// 	fmt.Println("Index: ", i, "value: ", v)
	// }

	i := 0
	for {
		if i >= len(a) {
			break
		}
		fmt.Println("Index:", i, "value:", a[i])
		i++
	}
}
