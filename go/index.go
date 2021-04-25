package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/cmplx"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
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
	strconvFunc()
	rangeFunc()
	rangeFunc2()
	mapsFunc()
	mapLiterals()
	mutatingMaps()
	functionValues()
	functionClosures()
	makeSlices2()
	byteFunc()
	variadicArgument()
	variadicArgument(10, 20, 30)
	variadicArgument(10, 20, 30, 40)
	s := []int{1, 2, 3}
	variadicArgument(s...)
	practice1()
	// LoggingSettings("test.log")
	// logging()
	// errorHandling()
	panicAndRecover()
	practice2()
	fibonacciFunc()
	methods()
	methodsFuncs()
	method()
	methodNonStruct()
	interfaceFunc()
	typeAssertion()
	stringer()
	customError()
	practice3()
	goroutine()
	channel()
}
func channel() {
	title := "channel"
	fmt.Println("---", title, "---")
	s1 := []int{1,2,3,4,5}
	s2 := []int{5,6,7,8,9}
	c := make(chan int)
	go goroutine2(s2, c)
	go goroutine1(s1, c)
	// blocking until x receives c
	// channel data structure: que
	x := <- c
	fmt.Println(x)
	y := <- c
	fmt.Println(y)
}
func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
func goroutine() {
	title := "goroutine"
	fmt.Println("---", title, "---")
	var wg sync.WaitGroup
	wg.Add(1)
	go normal("hello", &wg)
	normal2("world")
	wg.Wait()
}
func normal(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	// wg.Done()
}
func normal2(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func practice3() {
	title := "practice3"
	fmt.Println("---", title, "---")
	v := Vertex{3, 4}
	fmt.Println(v.Plus())
	fmt.Println(v)
}
func (v Vertex) Plus() int {
	return v.x + v.y
}
func (v Vertex) String() string {
	return fmt.Sprintf("X is %v Y is %v", v.x, v.y)
}
func customError() {
	title := "customError"
	fmt.Println("---", title, "---")
	if err := myCustomError(); err != nil {
		fmt.Println(err)
	}
}
func myCustomError() error {
	// Something wrong
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}
type UserNotFound struct {
	Username string
}
func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.Username)
}
func stringer() {
	title := "stringer"
	fmt.Println("---", title, "---")
	mike := Person2{"Mike", 39}
	fmt.Println(mike)
}
type Person2 struct {
	Name string
	Age int
}
// Use Stringer by String() when you want to modify the output of fmt.Println()
func (p Person2) String() string{
	return fmt.Sprintf("My name is %v.", p.Name)
}
func typeAssertion() {
	title := "typeAssertion"
	fmt.Println("---", title, "---")
	// var i interface{} = 10
	// fmt.Printf("%T %v\n", i, i)
	// do(i)
	do(10)
	do("Mike")
	do(true)

	var i int = 10
	ii := float64(10)
	fmt.Println(i, ii)
}
func do(i interface{}) {
	/*
	ii := i.(int)
	ii *= 2
	fmt.Println(ii)
	ss := i.(string)
	fmt.Println(ss + "!")
	*/
	switch v := i.(type) {
	case int :
		fmt.Println(v)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}
func interfaceFunc() {
	title := "interfaceFunc"
	fmt.Println("---", title, "---")
	var mike Human = &Person{"Mike"}
	var x Human = &Person{"X"}
	// var dog Dog = Dog{"dog"}
	// mike.Say()
	DriveCar(mike)
	DriveCar(x)
	// DriveCar(dog)
}
type Human interface {
	Say() string
}
type Person struct {
	Name string
}
type Dog struct {
	Name string
}
func (p *Person) Say() string{
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}
func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	}else{
		fmt.Println("Get out")
	}
}
func methodNonStruct() {
	title := "methodNonStruct"
	fmt.Println("---", title, "---")
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}
type MyInt int
func (i MyInt) Double() int {
	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", 1, 1)
	return int(i * 2)
}
func method() {
	title := "method"
	fmt.Println("---", title, "---")
	// v := Vertex{3, 4}
	// fmt.Println(Area(v))
	v := New(3, 4, 5)
	v.Scale3D(10)
	// fmt.Println(v.Area())
	fmt.Println(v.Area3D())
}
// this can be called within this package only
type Vertex struct{
	x, y int
}

func (v Vertex) Area() int{
	return v.x * v.y
}
func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}
// embedded
type Vertex3D struct{
	Vertex
	z int
}

func (v Vertex3D) Area3D() int{
	return v.x * v.y * v.z
}
func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}
// func Area(v Vertex) int {
// 	return v.x * v.y
// }
func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}
func methodsFuncs() {
	title := "methodsFuncs"
	fmt.Println("---", title, "---")
	v := Vertex2{3, 4}
	fmt.Println(Abs(v))
	fmt.Println(v)
}
func Abs(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func methods() {
	title := "methods"
	fmt.Println("---", title, "---")
	v := Vertex2{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(v)
}
type Vertex2 struct {
	X, Y float64
}
func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func fibonacciFunc() {
	title := "fibonacci"
	fmt.Println("---", title, "---")
	f := fibonacci()
	for i:= 0; i < 10; i++ {
		fmt.Println(f())
	}
}
func fibonacci() func() int {
	a, b := 1, 0
	return func() int {
		a, b = b, a+b
		return a
	}
}

func practice2() {
	title := "practice2"
	fmt.Println("---", title, "---")
	fmt.Println("Find the smallest num in the next slice")
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	smallest := l[0]
	for _, v := range l{
		if smallest > v {
			smallest = v
		}
	}
	fmt.Println(smallest)
	fmt.Println("How much is the total price?")
	m := map[string]int{
		"apple": 200,
		"grapes": 150,
    "orange": 80,
    "papaya": 500,
    "kiwi":   90,
	}
	sum := 0
	for _, v := range m{
		sum += v
	}
	fmt.Println(sum)
}
func panicAndRecover() {
	title := "panicAndRecover"
	fmt.Println("---", title, "---")
	save()
	fmt.Println("OK")
}
func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}
func thirdPartyConnectDB() {
	panic("Unable to connect database")
}
func errorHandling() {
	title := "errorHandling"
	fmt.Println("---", title, "---")
	file, err := os.Open("./index.go")
	if err != nil {
		log.Fatalln("Error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error!")
	}
	fmt.Println(count, string(data))

	if err = os.Chdir("test"); err != nil{
		log.Fatalln("Error")
	}
}
func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
func logging() {
	title := "logging"
	fmt.Println("---", title, "---")
	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	_, err := os.Open("fldja;")
	if err != nil{
		log.Fatalln("Exit", err)
	}

	log.Fatalf("%T %v", "error", "error")
	log.Fatalln("error!")

	fmt.Println("ok!")
}
func practice1() {
	title := "practice1"
	fmt.Println("---", title, "---")
	f := 1.11
	fmt.Println(int(f))
	s := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s[2:4])
	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi":30}
	fmt.Printf("%T %v\n", m, m)
}

func variadicArgument(params ...int) {
	title := "variadic argument"
	fmt.Println("---", title, "---")
	fmt.Println(len(params), params)
	for _, param := range params{
		fmt.Println(param)
	}
}

func byteFunc() {
	title := "byte"
	fmt.Println("---", title, "---")
	b := []byte{72, 73}
	fmt.Println(b)
	fmt.Println(string(b))
	
	c := []byte("hi")
	fmt.Println(c)
	fmt.Println(string(c))
}

func makeSlices2() {
	title := "make slices 2"
	fmt.Println("---", title, "---")
	c := make([]int, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

	d := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		d = append(d, i)
		fmt.Println(d)
	}
	fmt.Println(d)
}

func functionClosures() {
	title := "function closures"
	fmt.Println("---", title, "---")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func functionValues() {
	title := "function values"
	fmt.Println("---", title, "---")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(hypot)
	fmt.Println(compute(hypot))
	fmt.Println(math.Pow)
}
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func mutatingMaps() {
	title := "mutating maps"
	fmt.Println("---", title, "---")
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func mapLiterals() {
	title := "map literals"
	fmt.Println("---", title, "---")
	type Vertex struct {
		Lat, Long float64
	}
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

func mapsFunc() {
	title := "maps"
	fmt.Println("---", title, "---")
	type Vertex struct {
		Lat, Long float64
	}
	var m map[string]Vertex
	fmt.Println("The zero value of a map is nil. A nil map has no keys, nor can keys be added.")
	fmt.Println(m)
	
	m = make(map[string]Vertex)
	fmt.Println("The make function returns a map of the given type, initialized and ready for use.")
	fmt.Println(m)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}

func rangeFunc2() {
	title := "range2"
	fmt.Println("---", title, "---")
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func rangeFunc() {
	title := "range"
	fmt.Println("---", title, "---")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func strconvFunc() {
	title := "strconvFunc"
	fmt.Println("---", title, "---")
	// https://golang.org/pkg/strconv/
	i, _ := strconv.Atoi("14")
	fmt.Println(i)
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
