package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

var (
	i, j            int        = 1, 2
	ToBe            bool       = false
	MaxInt          uint64     = 1<<64 - 1
	z               complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Pi    = 3.14
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }
func one(x *int){
	*x = 1
}
func main() {
	forLoop()
	packages()
	imports()
	exportedNames()
	fmt.Println("The result of add1(1, 2) is ", add1(1, 2))
	fmt.Println("The result of add2(2, 3) is ", add2(2, 3))
	a, b := swap("hello", "world")
	fmt.Println("The result of a, b := swap('hello', 'world') -> ", a, b)
	/*
	func split(sum int) (x, y int) {
		x = sum *  4 / 9
		y = sum - x
		return // this is called 'naked return' dangerous!
	}
	*/
	d, e := split(9)
	fmt.Println("The result[0] of split(9) is ", d)
	fmt.Println("The result[1] of split(9) is ", e)
	variables()
	variablesWithInitializers()
	shortVariableDeclarations()
	basicTypes()
	zeroValues()
	
	var title string
	title = "Type conversions"
	fmt.Println("---------------------", title, "---------------------")
	x2, y2 := 3, 4
	fmt.Printf("Type of x2:%T, Type of y2:%T\n", x2, y2)
	f2 := math.Sqrt(float64(x2*x2 + y2*y2))
	fmt.Printf("Type of f2: %T\n", f2)
	z2 := uint(f2)
	fmt.Printf("Type of x2:%T, Type of y2:%T, Type of z2:%T\n", x2, y2, z2)

	title = "Type inference"
	fmt.Println("---------------------", title, "---------------------")
	v1 := 42
	v2 := "42"
	fmt.Printf("v1 is of type %T, v2 is of type %T\n", v1, v2)

	title = "Constants"
	fmt.Println("---------------------", title, "---------------------")
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	title = "Numeric Constants"
	fmt.Println("---------------------", title, "---------------------")
	// 	const (
		// 		Big = 1 << 100
		// 		Small = Big >> 99
		// 	)
		// func needInt(x int) int { return x*10 + 1}
		// func needFloat(x float64) float64 { return x * 0.1 }
		fmt.Println(needInt(Small))
		fmt.Println(needFloat(Small))
		fmt.Println(needFloat(Big))
		
	title = "pointer"
	fmt.Println("---------------------", title, "---------------------")
	var n int = 100
	one(&n) // passing address of n, and override the value from 100 to 1
	fmt.Println(n)
}

func forLoop() {
	var title string = "for"
	fmt.Println("---------------------", title, "---------------------")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Printf("i: %v -> sum: %v\n", i, sum)
	}
}

func packages() {
	var title string = "packages"
	fmt.Println("---------------------", title, "---------------------")
	fmt.Println("My favorite number is", rand.Intn(10))
}

func imports() {
	var title string
	title = "imports"
	fmt.Println("---------------------", title, "---------------------")
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(4))
}

func exportedNames() {
	title := "exported-names"
	fmt.Println("---------------------", title, "---------------------")
	// fmt.Println(math.pi)  // NG
	fmt.Println(math.Pi) // OK
}

func add1(x int, y int) int {
	title := "function add1"
	fmt.Println("---------------------", title, "---------------------")
	return x + y
}

func add2(x, y int) int {
	title := "function add2"
	fmt.Println("---------------------", title, "---------------------")
	return x + y
}

func swap(x, y string) (string, string) {
	title := "Multiple results"
	fmt.Println("---------------------", title, "---------------------")
	return y, x
}

func split(sum int) (x, y int) {
	title := "Named return values"
	fmt.Println("---------------------", title, "---------------------")
	x = sum * 4 / 9
	y = sum - x
	return // this is called 'naked return', and this is basically dangerous!
}

func variables() {
	title := "variables"
	fmt.Println("---------------------", title, "---------------------")
	var (
		c, python, java bool
		i 							int
	)
	fmt.Println("default values")
	fmt.Println(i, c, python, java)
}

func variablesWithInitializers() {
	title := "variables with initializers"
	fmt.Println("---------------------", title, "---------------------")
	var c, python, java = true, false, "no!"
	var i, j int = 10, 20
	fmt.Println("c: ", c, "\npython: ", python, "\njava: ", java, "\ni: ", i, "\nj: ", j)
}

func shortVariableDeclarations() {
	title := "Short variable declarations"
	fmt.Println("---------------------", title, "---------------------")
	k := 32
	js, solidity, golang := true, false, "go!"
	fmt.Println("k: ", k, "\njs: ", js, "\nsolidity: ", solidity, "\ngolang: ", golang)
}

func basicTypes() {
	title := "basic types"
	fmt.Println("---------------------", title, "---------------------")
	var (
		ToBe 		bool				= false
		MaxInt	uint64			= 1<<64 -1
		z				complex128	= cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("ToBe -> Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("MaxInt -> Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("z -> Type: %T Value: %v\n", z, z)
}

func zeroValues() {
	title := "zero values"
	fmt.Println("---------------------", title, "---------------------")
	var (
		i int
		f float64
		b bool
		s string
	)
	fmt.Printf("zero values of \nint: %v, \nfloat64: %v, \nbool: %v, \nstring: %v\n", i, f, b, s)
}