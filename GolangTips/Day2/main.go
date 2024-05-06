package main

import "fmt"

func main() {
	defer MultistageDefer()()

	fmt.Println("Main function called")
}
func MultistageDefer() func() {
	fmt.Println("Run initialization")
	return func() {
		fmt.Println("Run cleanup")
	}
}

//Run initialization
//Main function called
//Run cleanup
