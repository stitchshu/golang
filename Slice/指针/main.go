package main

import "fmt"

func main() {
	//slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//s1 := slice[2:5]
	//s2 := s1[2:6:7]
	//fmt.Println(s1)
	//fmt.Println(s2)

	s := []int{5}
	s = append(s, 7)
	s = append(s, 9)
	x := append(s, 11)
	y := append(s, 12)
	fmt.Println(s, x, y)
}
