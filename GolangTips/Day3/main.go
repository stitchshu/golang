package main

import "fmt"

func main() {
	a := make([]int, 5)
	a = append(a, 1)
	fmt.Println(a)
	b := make([]int, 0, 5)
	b = append(b, 1)
	fmt.Println(b)
	//[0 0 0 0 0 1]
	//[1]
}
