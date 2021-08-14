GO
---

## Why Go

- Compiles to a single binary file (compared to node.js which is interpreted language, hence faster)
- No runtimes to worry about ( once compiled to a binary file, no need to worry about the runtime)
- Statically typed, so no surprises in runtime
- Sort of object oriented
- Concurrency
- Cross platform
- Built-in testing
- package management

<br/>

## GO CLI

|Command| Description
|:----------|:----------|
| go build | Compiles a bunch of source code and generate an executable file|
| go run | Compiles and executes one or many files|
|go fmt | formats all code in each file in the current directory |
|go install | Compiles and install a package |
| go get | downloads raw source code of package|
| go test | runs any test associated with the current project |
-----

<br/>

## Go Packages

<br/>

`Package  == Project == workspace`

Package keyword is used to define the workspace of the current project. Package can have multiple files. The main requirement is to define the package name on every file where it belongs to.

In Go we have 2 types of packages:

1. executable: Generates a file that we can run as executable with `go build ...` to accomplish tasks

2. Reusable: libraries/ code dependencies. Used as helpers (common functions for reusable code). These packages are not executable.

The name of the package decides whether we are making reusable or executable package.

Package 'main' is used to make executable type package

```
Package main --> go build --> main.exe (or main)

Package package1 --> go build --> nothing
```

|Package| Definition|
|---|---|
|main ( Executable)| Defines a package that can be compiledand then executed. **Must have a func called 'main'**
|calculator (Reusable)| Defines a package that can be  used as dependency (helper code)|
|uploader (Reusable)| Defines a package that can be used as dependency (helper code)|

<br/>

## Import Statements

Import statement is used to give our package access to some code, written inside another package.

``` 
import "fmt" // short form of format

import (    // multiple imports
	"fmt"
	"math"
)
```
This means giving our package "main" access to all code and functions exported by package "fmt". 
"fmt" is a standard library included inside go programming language by default. "fmt" library is used to print-out a lot of different information in the terminal

In Go, a name is exported if it begins with a capital letter. 
For example, Println is exported from the fmt package
fmt.<ins>Println</ins>("")

<br/>

## How to organize the file in main.go?

```
// package declaration
    package main  

// import other packages that we need
    import "fmt" 

// declare functions
    func main() {
        
    }
```

<br/>

## Variable Declaration

```
var card string = "abc"
card := "abc"

var --> create a new variable
card --> The name of the variable is card
string --> Only a value of type "string" will be assigned to this variable
"abc" --> Assign the value abc to the variable card

:= helps "Go" to infer the type of variable to define (available on right-side of =).
This should be used only when we are defining a new variable. During reassignment this is not required.
card := "abc" // initialization
card = "a"  // reassignment
```

Go is a static type language.

Basic Types

1. bool (true,false)
2. string ("abc")
3. int (1,10,45)
4. float64 (10.00001)
5. rune (represents a unicode point, normally used fdor character code)


A var statement can be at package or function level.
```
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```
- A var declaration can include initializers, one per variable. If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

- Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

- Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

- Variables declared without an explicit initial value are given their zero value.

    The zero value is:

    * 0 for numeric types
    * false for the boolean type, and
    * "" (the empty string) for strings.


**Constants**
Constants are declared with const keyword

Constants can be character, string, boolean, or numeric values.Constants cannot be declared using the := syntax.

```
const Truth = true
Truth = false // not possible
fmt.Println("Go rules?", Truth)
```
<br/>

## Defer
A defer statement defers the execution of a function until the surrounding function returns.

```
package main

import "fmt"

func main() {
	fmt.Println("hey")
	defer fmt.Println("world")
	defer fmt.Println("world1")
fmt.Println("hi")
	fmt.Println("hello")
}

output:
hey
hi
hello
world1
world
```
Deferred function calls are pushed onto a `stack`. When a function returns, its deferred calls are executed in `last-in-first-out order`.

## Type conversions

The expression *T(v)* converts the value v to the type T.

```
i := 42
f := float64(i)  // T = float64, v = i
u := uint(f)    // T = uint, v = i
```

## Functions and Return Types

```
    func newCard() string {
        return "abcd"
    }

    func --> define a function
    newCard --> name of the function
    string --> return type of the function
```

Files in the same package can freely call functions defined in other files.

```
** main.go **
------------

package main
 
func main() {
    printState()
}

** state.go **
---------------

package main
 
import "fmt"
 
func printState() {
    fmt.Println("California")
}

```

Now use `go run main.go state.go` to run the above example. 
Notice main.go is using the function defined in state.go without importing any pkg.

A function can take zero or more arguments.

For example

```
    func add(x int, y int) int {
        return x+y;
    }
```

When two or more consecutive named function parameters share a type, we can omit the type from all but the last.
In the above example, we shortened

```
    x int, y int

    to

    x, y int


    func add(x, y int) int {
        return x+y;
    }

```

A function can return any number of results.
The below swap function returns two strings

```
    func swap(x,y string) (string, string) {
        return y, x
    }

    func main() {
        a,b := swap("hello", "world")
        fmt.Println(a, b)
    }
```

Go's return values may be named. If so, they are treated as variables defined at the top of the function. These names should be used to document the meaning of the return values. A return statement without arguments returns the named return values. This is known as a "naked" return.
Naked return statements should be used only in short functions, else they can harm readability in longer functions.

```
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17)) // prints 7 10
}
```


## Array and Slice

array: Fixed length list of elements
slice: An array that can grow or shrink (dynamic array)

A slice or array must be defined with a data-type and every element must be of same type.

**Arrays**

The type [n]T is an array of n values of type T.

```
var a [10]int

declares a variable a as an array of ten integers.
```
An array's length is part of its type, so arrays cannot be resized.

**Slice**
A slice does not store any data, it just describes a section of an underlying array. 

A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.


```
cards := []string{"Ace of diamonds"}  // initialize a slice
cards = append(cards, "Six of Spades")

[]string --> define a slice of type string
string{"Ace of diamonds"} --> initialize the slice of first element
append --> function to add element at the end of the slice, this returns a new slice and do not modify the old cards variable

```
When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

```
fruits : []string{"apple", "orange", "banana", "grape"}
fruits[0] --> "apple"

// fruits[startIndexInclusive:endIndexExclusive]
fruits[0:2] --> ["apple", "orange"]
fruits[:2] --> ["apple", "orange"]
fruits[2:] --> ["banana", "grape"]
```

A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
The zero value of a slice is nil.

A nil slice has a length and capacity of 0 and has no underlying array.

**Creating a slice with make**
Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:

```
a := make([]int, 5)  // len(a)=5
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
```

**Byte Slice** []byte represents the ASCII char of a string.



When we create a slice, internally "Go" will create 2 data-structures
1. A structure that records the length of the slice, the capacity of the slice and the reference to the underlying array
2. Array, which represents all the items. "pointer to head" mentioned above points to Array.

|Address | Value |
|---|---|
| 0001 | slice({length, capacity, point to head})|
| 0002 | array([]string{"a", "b", "c", "d"})

mySlice in the below example refers to slice present in memory 0001.

```
func main() {
	mySlice := []string{"Hi", "There", "How", "are", "you"}
	fmt.Println(mySlice)
}
```

Update a slice.

```
func main() {
	mySlice := []string{"Hi", "There", "How", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}
func updateSlice(s []string){
	s[0] = "Bye"
}  
```
Go will create a new slice data-structure for the updateSlice, but the new slice will point to the same array underneath.

## Range

The range form of the for loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
You can skip the index or value by assigning to _.

```
for i, _ := range pow
for _, value := range pow
```

If we only want the index, you can omit the second variable.
```
for i := range pow
```

## For loop

Go has only one looping construct, the for loop.

The basic for loop has three components separated by semicolons:

* the init statement: executed before the first iteration
* the condition expression: evaluated before every iteration
* the post statement: executed at the end of every iteration

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

```
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

```
The init and post statements are optional.

```
package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
```

We can also drop the initial semicolon. The <ins>for loop</ins> will behave as while loop (present in javascript)

```
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```

Infinite loop:

```
package main

func main() {
	for {
	}
}
```

Example for slice. **range** is the keyword we use when we try to iterate over every record inside a slice

```
for index,value := range of values {

}

* index --> index of the current element
* value --> value of the current element we are iterating over
* range values --> "Go" interpret that there is a slice of records, it will take the slice and iterate over every element inside of it.
```


## OO Approach vs Go Approach

Go is not an object-oriented language, so no concept of classes in Go.

|Object-oriented|    Go  |
|:---:|:---|
|class Deck{}| new type Deck extends base type|
| define attributes and methods to the class|  function with 'deck' as receiver. A function with a receiver is like a method|
---------

Since we are creating a custom type 'deck', it gives us the ability to create custom functions that only works with the 'deck' type

```
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

(d deck) --> referred to as a receiver on a function print
```

Any variable of type "deck" can now get access to the "print" method. Receiver sets up method on the variable that we create.
By convention, we usually refer to the receiver with a one (or two) letter abbreviation that matches the type of the receiver.

Whenever there is a variable we do not want to use, replace the same with underscore (_).

for example:
```
func print() {
	for _, card := range cards {
		fmt.Println(card)
	}
}
```
In the above example we are printing the value of card, we do not need the index, so we replaced index with _.


## Testing in Go

To make a test, create a new file ending with `_test.go`

to run all testsin a package use `go test`

`go env -w GO111MODULE=auto` decides where to run the go commands inside/outside GoPath

To test functions in Go we have to take care of cleanup manually.



## Structs 

data-structure in go. Collection of different properties that are related together.

3 ways of defining structs

```
type person struct{
    firstName string
    lastName string
}


1. alex := person{"Alex", "Anderson"}   
// Problem with the above approach is the order should 
// be same as the order defined while creating 
// the struct person

2. alex := person{firstName: "Alex", lastName: "Anderson"}

3.	var alex1 person // alex has no property assigned to it. Go assigns "zero value" to each of the fields defined in the struct
	alex1.firstName = "Alex"
	alex1.lastName = "Anderson"
	fmt.Println(alex1)
	fmt.Printf("%+v", alex1)
```
%c --> character
%d --> integer

**Reference of one struct into another**
```
type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo 
}

"contact contactInfo" can be replaced with simply contactInfo, where Go will see the struct and will infer the field as "contactInfo"    and type contactInfo

```


## Pointers

A pointer holds the memory address of a value.
Go is "passByValue" language. 
Whenever we try to update a struct, "Go" will get the value of the struct and copy the value in a different location of memory. This means whenever we are changing struct, it is creating another copy of the same value rather than referencing to the same memory location

```
jim := person{
    firstName: "Jim",
    lastName:  "Anderson",
    contact: contactInfo{
        email:   "jim@gmail.com",
        zipCode: 560048,
    },
}
jimPointer := &jim
jimPointer.updateName("Jimmy")

func (p *person) updateName(newName string) {
    (*p).firstName = newName
}

*person --> This is a type description. In the above example pointer to a person Jim

*p --> This is an operator. It means we want to manipulate 

```

& variable : fetch the memory address of the value this variable is pointing at 

*pointer: fetch the value this memory address is pointing at.

|Address| value |
|---|---|
|0001| person{firstname:"Jim"}|


* Turn `address` into `value` with `*address`
* Turn `value` into `address` with `&value`


**Short-cut**
```
# first example
jimPointer := &jim
jimPointer.updateName("Jimmy")

func (p *person) updateName(newName string) {
    (*p).firstName = newName
}

- jimPointer --> pointer to a person
- *person --> pointer to a person

# change the above example to:
jim.updateName("Jimmy")

func (p *person) updateName(newName string) {
    (*p).firstName = newName
}

- jim --> type of person

```

In the second example, "Go" will check jim as type of person and *person as pointer to a person and it will automatically turn variable (type of person) into adress(pointer to a person).

## Value and Reference Type

| Value Types | Reference Types |
|---|---|
| int | slices |
|float| maps |
|string| channels |
|bool| pointers |
|struct| functions |


For value types use `pointers` to change the value in a function
For reference type, pointers not required

## Map

Maps are key-value pairs. In map the key-value pairs are statically typed. All keys must of of same type. All values must be of same type.

One way to declare a map 
```
colors := map[string]string{
    "red": "#ff000",
    "white": "#ffffff",
    "black": "#000000",
}
declares a mao where all keys are string and all values are string
```

Another ways to declare a map
```
var colors map[string]string
        OR
colors := make(map[string]string)

```
add values to Map, Always use [] brackets.
```
colors["white"] = "#fff"
```

Delete values from Map
```
delete(colors, "white")
```
Looping over Maps

```
for color, hex := range colors {
    fmt.Println("Hex code for", color, "is", hex)
}
```

## Maps vs structs

|Map | Struct|
|:---|:---|
| All keys must be of same type | keys are not stringly typed|
| All values must be of same type | Values can be of different type |
| Keys are indexed - we can iterate over them | Keys don't support indexing|
| Reference type (pass by reference) | Value type (pass by value)|
|Don't need to know all keys at compile time (Schemaless)| Need to know all different fields in compile time (Schema is required)|


## Interface

An interface type is defined as a set of method signatures.
A value of interface type can hold any value that implements those methods.
A type implements an interface by implementing its methods. 

Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

```
(value, type)
```

An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type.

Considering below example:

	```
	type bot interface {
		getGreeting() string
	}
	type englishBot struct{}
	type spanishBot struct{}

	func main() {
		eb := englishBot{}
		sp := spanishBot{}

		printGreeting(eb)
		printGreeting(sp)
	}
	func printGreeting(b bot) {
		fmt.Println(b.getGreeting())
	}

	func (englishBot) getGreeting() string {
		// very custom logic for generation english greeting
		return "Hello!"
	}

	func (spanishBot) getGreeting() string {
		return "Hola!"
	}
	```
	Any type (englishBot or spanishBot) available on the program with a receiver function (getGreeting) and the return type of the receiver function matches the function defined in interface (in this case string), then the type is a member of the interface type (bot). We defined functions inside interfaces here.

	We can define multiple arg list and return types to the function defined inside interface like:

	```
	type bot interface {
		getGreeting(string, int) (string, error)
	}
	```

	We can also define multiple functions inside an interface. Remember there are no commas in between each func

	```
	type bot interface {
		getGreeting(string, int) (string, error)
		getBotVersion() float64
		respondToUser(user) string
	}
	```

| Concrete type | Interface Type |
|---|---|
|map, struct, int, string, englishBot | bot|

We can create a value directly with Concrete type like map, struct,etc., but we can not create a value directly from interface.

* Interfaces are not generic types (Go does not  have generic types)
* Interfaces are implicit : We don't manually have to say that our custom type satisfies some interfaces.
* Interfaces are contract to help us manage types : GARBAGE IN -> GARBAGE out. If our custom type implementation of a function is broken the interfaces won't help us.
* We can wrap interface inside another interface : We can assemble multiple interfaces into one interface

Check the below example for http package

```
type response struct{
	Status string
	StatusCode int
	Body io.ReadCloser
}

type ReadCloser interface {
	Reader
	Closer
}

type Reader interface {
	Read([]byte) (int, error)
}

type Closer interface {
	Close() error
}
```

The value Body can be value of any type which satisfies ReadCloser interface.


## Channels and Go Routines

A goroutine is a lightweight thread managed by the Go runtime. Goroutines run in the same address space, so access to shared memory must be synchronized.

Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`.

```
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.

ch := make(chan int)
```
By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.


Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

```
ch := make(chan int, 100)
```
Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.


A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

```
v, ok := <-ch
```
**Note**: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

The select statement lets a goroutine wait on multiple communication operations.

A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

```
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

Consider the below example:

```
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		return
	}
	fmt.Println(link, "is up")
}
```

The flow of the above code is as below:
1. Take the first link
2. Make http request
3. Wait for the response, log it
4. Take the second link and so on

Thus, we ended up in a blocking code.
Use Go Routings to run each of these requests in parallel

When we launch a go program, we automatically create a Go routine. This Go routine takes every line of code  inside the program and executes them one by one.

The code can be changed into non-blocking by adding a "go" keyword in front of the function inside for loop

```
for _, link := range links {
	go checkLink(link)
}
```
Child routines are created by "go" keyword.

"Go Scheduler" works with one-CPU core per machine. Scheduler runs one go routine until it finishes or makes a blocking call (like HTTP request). Then it pauses the go routine and send the control back to the main routine.
For multiple-core, go scheduler runs one thread on each "logical core". By default, Go tries to use one-CPU core.

**Concurrency is not parallelism**

|Concurrency|Parallelism|
|---|----|
| A program is concurrent if it has the ability to pick up multiple go routines at a time. All these Go routines must work in a single-core. Thus, we can do multiple things by context switching | Parallelism is when we include multiple core CPU in our machine. With Parallelism we are saying we can multiple things at the same time.
|


Channels are used for communication between different running Go routines. Channel can have a specific type. We would not be able to work with channels with type string and try to send int or float to it.

```
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- "might be down"
		return
	}
	fmt.Println(link, "is up")
	c <- "Yep, it's up"
}
```
The program starts execution. But after the first link, the program exited.
WHen Go has spawned all go routines from main routine 
```	
for _, link := range links {
	go checkLink(link, c)
}
```
The main routine will wait for the channel to receive some information
```
fmt.Println(<-c)
```
This is a blocking code. Receiving message from a channel is a blocking code.
To improve this code use:

```
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
```

Another way to use for loop for channel is

```
for l := range c {
	go checkLink(l, c)
}
```
where we are saying wait for the channel to return some value. After it returns value assign it to l.

**Sleeping a Go routing**

package time.Sleep pauses the current goroutine for least d duration.


## Function values

functions can be passed as arguments to another function.( similar to JS)

