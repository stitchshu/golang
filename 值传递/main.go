package main

import "fmt"

type Man struct{}

type Woman struct{}

func (*Man) Say() {
	fmt.Println("Man Say")
}
func (Woman) Say() {
	fmt.Println("woman say")
}
func main() {
	man := Man{}
	man.Say()
	woman := Woman{}
	woman.Say()
	ptrWoman := &Woman{}
	ptrWoman.Say()
}
