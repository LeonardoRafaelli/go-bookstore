package main

import "fmt"

func main() {
	login()
}

func login() {

	fmt.Println(4 >> 2)
	defer deferTry()

	fmt.Println("Chegou aqui")
	fmt.Println("Chegou aqui X2")
}

func deferTry() {
	fmt.Println("Chegou X3")
}
