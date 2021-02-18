package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

var (
	c, python, java bool
	i, j            int        = 1, 2
	ToBe            bool       = false
	MaxInt          uint64     = 1<<64 - 1
	z               complex128 = cmplx.Sqrt(-5 + 12i)
)

func add1(x int, y int) int {
	return x + y
}
func add2(x, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
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
	fmt.Println(math.Pi) // OK
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
	/*
		func split(sum int) (x, y int) {
			x = sum *  4 / 9
			y = sum - x
			return // this is called 'naked return' dangerous!
		}
	*/
	fmt.Println(split(17))

	title = "Variables"
	fmt.Println("---------------------", title, "---------------------")
	// var c, python, java bool // declared at the top of this code
	var i int
	fmt.Println(i, c, python, java)

	title = "Variables with initializers"
	fmt.Println("---------------------", title, "---------------------")
	// var i, j int = 1, 2 // declred at the top of this code
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	title = "Short variable declarations"
	fmt.Println("---------------------", title, "---------------------")
	k := 3
	javascript, solidity, golang := true, false, "no!"
	fmt.Println(k, javascript, solidity, golang)

	title = "Basic types"
	fmt.Println("---------------------", title, "---------------------")
	// var (
	// 	ToBe            bool       = false
	// 	MaxInt          uint64     = 1<<64 - 1
	// 	z               complex128 = cmplx.Sqrt(-5 + 12i)
	// )
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	var x uint64 = 1
	fmt.Println(x<<64-1)
	

	title = ""
	fmt.Println("---------------------", title, "---------------------")

	title = ""
	fmt.Println("---------------------", title, "---------------------")

	title = ""
	fmt.Println("---------------------", title, "---------------------")

}
