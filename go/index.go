package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}
func main() {
	var title string = "packages"
	fmt.Println("---------------------", title, "---------------------")


	fmt.Println("My favorite number is", rand.Intn(10))
	
	title = "imports"
	fmt.Println("---------------------", title, "---------------------")
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
	
	title = "exported-names"
	fmt.Println("---------------------", title, "---------------------")
	// fmt.Println(math.pi)  // NG
	fmt.Println(math.Pi)  // OK
	title = "functions"
	fmt.Println("---------------------", title, "---------------------")
	fmt.Println("The result of add(1, 2) is ", add(1, 2))
	title = ""
	fmt.Println("---------------------", title, "---------------------")


	title = ""
	fmt.Println("---------------------", title, "---------------------")


	title = ""
	fmt.Println("---------------------", title, "---------------------")


	title = ""
	fmt.Println("---------------------", title, "---------------------")


	title = ""
	fmt.Println("---------------------", title, "---------------------")


	title = ""
	fmt.Println("---------------------", title, "---------------------")


	title = ""
	fmt.Println("---------------------", title, "---------------------")


	title = ""
	fmt.Println("---------------------", title, "---------------------")


	title = ""
	fmt.Println("---------------------", title, "---------------------")


	title = ""
	fmt.Println("---------------------", title, "---------------------")


	
	
	
}
