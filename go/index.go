package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
	"strings"
)
func init() {
	fmt.Println("Hello Go World!")
}
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
	forLoop()
	forContinued()
	while()
	forever()
	ifFunc()
	ifWithShortStatement()
	ifAndElse()
	switchFunc()
	switchEvaluationOrder()
	switchWithoutCondition()
	deferFunc()
	deferMult()
	pointer()
	pointer2()
	structs()
	structsFields()
	structsPointers()
	structsLiterals()
	arrayFunc()
	slices()
	slicesPointers()
	sliceLiterals()
	sliceBounds()
	sliceLenCap()
	nilSlices()
	makingSlices()
	slicesOfSlice()
	appendFunc()
}

func appendFunc() {
	title := "appendFunc"
	fmt.Println("---", title, "---")
	var s []int
	printSlice(s)
	
	fmt.Println("append works on nil slices.")
	s = append(s, 0)
	printSlice(s)
	
	fmt.Println("The slice grows as needed.")
	s = append(s, 1)
	printSlice(s)

	fmt.Println("We can add more than one element at a time.")
	s = append(s, 2, 3, 4, 5)
	printSlice(s)	
}

func slicesOfSlice() {
	title := "slicesOfSlice"
	fmt.Println("---", title, "---")
	fmt.Println("Create a tic-tac-toe board.")
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i:= 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func makingSlices() {
	title := "makingSlices"
	fmt.Println("---", title, "---")
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("a", a)

	c := b[:2]
	printSlice2("c", c)

	d := c[2: 5]
	printSlice2("d", d)
}
func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func nilSlices() {
	title := "nilSlices"
	fmt.Println("---", title, "---")
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func sliceLenCap() {
	title := "sliceLenCap"
	fmt.Println("---", title, "---")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	fmt.Println("Slice the slice to give it zero length")
	s = s[:0]
	printSlice(s)

	fmt.Println("Extend its length")
	s = s[:4]
	printSlice(s)
	
	fmt.Println("Drop its first two values")
	s = s[2:]
	printSlice(s)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func sliceBounds() {
	title := "sliceBounds"
	fmt.Println("---", title, "---")
	s := []int{2, 3, 5, 7, 11, 13}
	s =  s[1:4]
	fmt.Println(s) // 3, 5, 7
	s = s[:2]
	fmt.Println(s) // 3, 5
	s = s[1:]
	fmt.Println(s) // 5
}
func sliceLiterals() {
	title := "sliceLiterals"
	fmt.Println("---", title, "---")
	fmt.Println(`A slice literal is like an array literal without the length.`)
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("q: ", q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println("r: ", r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("s: ", s)
}
func slicesPointers() {
	title := "slicesPointers"
	fmt.Println("---", title, "---")
	fmt.Println(`A slice does not store any data, it just describes a section of an underlying array.`)
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]		// john, paul
	b := names[1:3]		// paul, george
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)	// [John Paul] [XXX George]
	fmt.Println(names)// [John Paul George Ringo]
}
func slices() {
	title := "slices"
	fmt.Println("---", title, "---")
	fmt.Println(`An array has a fixed size.
	A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
	In practice, slices are much more common than arrays.`)
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // is slicing [3, 5, 7]
	fmt.Println(s)

	oneToFive := []int{1, 2, 3, 4, 5}
	fmt.Println(oneToFive)
	fmt.Println(oneToFive[2:3])	
	fmt.Println(oneToFive[2:3][0])	
}
	func arrayFunc() {
	title := "arrayFunc"
	fmt.Println("---", title, "---")
	fmt.Println("An array's length is part of its type, so arrays cannot be resized.")
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println("a[0]: ", a[0])
	fmt.Println("a[1]: ", a[1])
	fmt.Println("a: ", a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("primes: ", primes)
}
func structsLiterals() {
	title := "structsLiterals"
	fmt.Println("---", title, "---")
	type Vertex struct {
		X, Y int
	}
	var (
		v1 = Vertex{1, 2}		// has type Vertex
		v2 = Vertex{X: 1}		// Y: 0 is implicit
		v3 = Vertex{}				// X: 0, Y: 0
		p  = &Vertex{1, 2}	// has type *Vertex
	)
	fmt.Println("v1: ", v1)
	fmt.Println("v2: ", v2)
	fmt.Println("v3: ", v3)
	fmt.Println("p: ", p)
	p.X = 1e3
	fmt.Println("p.X: ", p.X)
}
func structsPointers() {
	title := "structsPointers"
	fmt.Println("---", title, "---")
	type Vertex struct {
		X int
		Y int
	}
	v := Vertex{1, 2}
	fmt.Println(v)
	p := &v
	/*
	1e□□という数値の意味
	大きい数字は 1e3 のように e を使って表すことがあります。e は 10 のべき乗を表します。
	例えば、1e3 は 10^3=1000 を表します。1e9=1000000000
	*/
	p.X = 1e9
	fmt.Println(v)
}
func structsFields() {
	title := "structsFields"
	fmt.Println("---", title, "---")
	type Vertex struct {
		X int
		Y int
	}
	v := Vertex{1, 2}
	fmt.Println("v.X: ", v.X)
	v.X = 4
	fmt.Println("v.X: ", v.X)
	fmt.Println("v.Y: ", v.Y)
}
func structs() {
	title := "structs"
	fmt.Println("---", title, "---")
	type Vertex struct {
		X int
		Y int
	}
	fmt.Println(Vertex{1, 2})
}
func pointer2() {
	title := "pointer2"
	fmt.Println("---", title, "---")
	i, j := 42, 2701

	iAddr := &i												// point to i
	fmt.Println("iAddr: ", iAddr)			// see the address of i
	fmt.Println("*iAddr: ", *iAddr)		// read i through the pointer
	*iAddr = 21												// set i through the pointer
	fmt.Println("*iAddr: ", *iAddr)		// read i through the pointer
	fmt.Println("i: ", i)							// see the new value of i

	jAddr := &j												// point to j
	fmt.Println("j: ", j)
	fmt.Println("jAddr: ", jAddr)
	*jAddr = *jAddr / 37							// divide j through the pointer
	fmt.Println(j)										// see the new value of j
}
func deferMult() {
	title := "deferMult"
	fmt.Println("---", title, "---")
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)

	}
	fmt.Println("done")
}
func deferFunc() {
	title := "deferFunc"
	fmt.Println("---", title, "---")
	defer fmt.Print(" later.")
	defer fmt.Print("postponed something")
	defer fmt.Print(" that being ")
	defer fmt.Print("Defer means")
	defer fmt.Println("!!")
	defer fmt.Print("World ")

	fmt.Print("Hello ")
}
func switchWithoutCondition() {
	title := "switchWithoutCondition"
	fmt.Println("---", title, "---")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
func switchEvaluationOrder() {
	title := "switchEvaluationOrder"
	fmt.Println("---", title, "---")
	fmt.Println("When's Saturday?")
	now := time.Now()
	fmt.Println("now: ", now)
	today := time.Now().Weekday()
	fmt.Println("today: ", today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In 2 days")
	default:
		fmt.Println("Too far away...")
	}
}
func switchFunc() {
	title := "switch"
	fmt.Println("---", title, "---")
	// fmt.Print("Go runs on ")
	os := runtime.GOOS;
	fmt.Println("os; ", os)
	switch os {
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
