package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add1(x int, y int) int {
	return x + y
}
func add2(x, y int) int {
	return x + y
}
func swap(x, y string)(string, string) {
	return y, x
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
	fmt.Println("The result of add(1, 2) is ", add1(1, 2))
	title = "Functions continued"
	fmt.Println("---------------------", title, "---------------------")
	fmt.Println("The result of add(1, 2) is ", add2(1, 2))

	title = "Multiple results"
	fmt.Println("---------------------", title, "---------------------")
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	title = "Named return values"
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
