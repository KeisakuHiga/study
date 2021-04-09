package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
)

func main() {
	packages()
	imports()
	exportedNames()
	fmt.Println("The result of add1(1, 2) is ", add1(1, 2))
	fmt.Println("The result of add2(2, 3) is ", add2(2, 3))
	a, b := swap("hello", "world")
	fmt.Println("The result of a, b := swap('hello', 'world') -> ", a, b)
	d, e := split(9)
	fmt.Println("The result[0] of split(9) is ", d)
	fmt.Println("The result[1] of split(9) is ", e)
	variables()
	variablesWithInitializers()
	shortVariableDeclarations()
	basicTypes()
	zeroValues()
	typeConversions()	
	typeInference()
	constants()
	numericConstants()
	pointer()
	forLoop()
	forContinued()
	while()
	forever()
	ifFunc()
	ifWithShortStatement()
	ifAndElse()
	switchFunc()
}

func switchFunc() {
	title := "switch"
	fmt.Println("---", title, "---")
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// feedsb, opendsb,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func ifAndElse() {
	title := "if and else"
	fmt.Println("---", title, "---")
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
		} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
func ifWithShortStatement() {
	title := "if with a short statement"
	fmt.Println("---", title, "---")
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func ifFunc() {
	title := "if"
	fmt.Println("---", title, "---")
	fmt.Println("sqrt(2): ", sqrt(2))
	fmt.Println("sqrt(-4)", sqrt(-4))
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func forever() {
	title := "forever"
	fmt.Println("---", title, "---")
	// for {}
}

func while() {
	title := "while"
	fmt.Println("---", title, "---")
	sum := 3
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}
}

func forContinued() {
	var title string = "for continued"
	fmt.Println("---", title, "---")
	sum := 1
	for ; sum < 1000; {
		fmt.Println(sum)
		sum += sum
	}
}

func forLoop() {
	var title string = "for"
	fmt.Println("---", title, "---")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Printf("i: %v -> sum: %v\n", i, sum)
	}
}

func packages() {
	var title string = "packages"
	fmt.Println("---", title, "---")
	fmt.Println("My favorite number is", rand.Intn(10))
}

func imports() {
	var title string
	title = "imports"
	fmt.Println("---", title, "---")
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(4))
}

func exportedNames() {
	title := "exported-names"
	fmt.Println("---", title, "---")
	// fmt.Println(math.pi)  // NG
	fmt.Println(math.Pi) // OK
}

func add1(x int, y int) int {
	title := "function add1"
	fmt.Println("---", title, "---")
	return x + y
}

func add2(x, y int) int {
	title := "function add2"
	fmt.Println("---", title, "---")
	return x + y
}

func swap(x, y string) (string, string) {
	title := "Multiple results"
	fmt.Println("---", title, "---")
	return y, x
}

func split(sum int) (x, y int) {
	title := "Named return values"
	fmt.Println("---", title, "---")
	x = sum * 4 / 9
	y = sum - x
	return // this is called 'naked return', and this is basically dangerous!
}

func variables() {
	title := "variables"
	fmt.Println("---", title, "---")
	var (
		c, python, java bool
		i 							int
	)
	fmt.Println("default values")
	fmt.Println(i, c, python, java)
}

func variablesWithInitializers() {
	title := "variables with initializers"
	fmt.Println("---", title, "---")
	var c, python, java = true, false, "no!"
	var i, j int = 10, 20
	fmt.Println("c: ", c, "\npython: ", python, "\njava: ", java, "\ni: ", i, "\nj: ", j)
}

func shortVariableDeclarations() {
	title := "Short variable declarations"
	fmt.Println("---", title, "---")
	k := 32
	js, solidity, golang := true, false, "go!"
	fmt.Println("k: ", k, "\njs: ", js, "\nsolidity: ", solidity, "\ngolang: ", golang)
}

func basicTypes() {
	title := "basic types"
	fmt.Println("---", title, "---")
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
	fmt.Println("---", title, "---")
	var (
		i int
		f float64
		b bool
		s string
	)
	fmt.Printf("zero values of \nint: %v, \nfloat64: %v, \nbool: %v, \nstring: %v\n", i, f, b, s)
}

func typeConversions() {
	title := "Type conversions"
	fmt.Println("---", title, "---")
	x, y := 3, 4
	fmt.Printf("Type of x: %T, Type of y: %T\n", x, y)

	f := math.Sqrt(float64(x*x + y*y))
	fmt.Printf("Type of f: %T\n", f)

	z := uint(f)
	fmt.Printf("Type of z: %T\n", z)
}

func typeInference() {
	title := "Type inference"
	fmt.Println("---", title, "---")
	v1 := 42
	v2 := "42"
	fmt.Printf("v1 is of type %T, v2 is of Type %T\n", v1, v2)
}

func constants() {
	title := "Constants"
	fmt.Println("---", title, "---")
	const World = "世界"
	fmt.Println("Hello", World)
	const (
		Pi		= 3.14
	)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)
}

func numericConstants() {
	title := "Numeric Constants"
	fmt.Println("---", title, "---")
	const (
		Big = 1 << 100
		Small = Big >> 99
	)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
func needInt(x int) int { return x*10 + 1}
func needFloat(x float64) float64 { return x * 0.1}

func pointer() {
	title := "pointer"
	fmt.Println("---", title, "---")
	var n int = 100
	fmt.Println("print value of n -> ", n)
	fmt.Println("print address of n -> ", &n)
	fmt.Println("function one(x int) overrides the arguments' value to 1")
	one(&n)
	fmt.Println("print value of n -> ", n)
	fmt.Println("print address of n -> ", &n)
}
func one(x *int){
	*x = 1
}
