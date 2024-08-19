package main

import "fmt"

func main() {
	doSomething()
	fmt.Println("After doSomething")
}

func doSomething(){
	defer letsRecover()
	fmt.Println("1")
	panic(doSomething)
	fmt.Println("2")
}

func letsRecover(){
	recover()
}