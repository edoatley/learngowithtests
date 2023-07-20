# learngowithtests

- [learngowithtests](#learngowithtests)
  - [Running code](#running-code)
  - [Creating a module](#creating-a-module)
  - [Writing tests](#writing-tests)
  - [Go doc](#go-doc)
  - [Examples](#examples)
  - [Benchmarking](#benchmarking)
  - [Test Coverage](#test-coverage)
  - [Interesting Language notes](#interesting-language-notes)
    - [Arrays and slices](#arrays-and-slices)
    - [First class functions](#first-class-functions)
    - [Interfaces](#interfaces)
    - [Pointers](#pointers)
  - [Maps](#maps)
  - [Cannot import main: a Go Module gotcha](#cannot-import-main-a-go-module-gotcha)


Following along with https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

## Running code

You can run the file directly with `go run hello.go`
After creating a module (see below) you can run `go run .` from the root of the module

## Creating a module

Running `go mod init SOMENAME` will create a module with the name `SOMENAME` and you can then run

## Writing tests

Writing a test is just like writing a function, with a few rules

- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only t *testing.T
- In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file

You can run the tests by running `go test`

## Go doc
Another quality of life feature of Go is the documentation. You can launch the docs locally by running `godoc -http :8000`. 
If you go to localhost:8000/pkg you will see all the packages installed on your system.

The vast majority of the standard library has excellent documentation with examples. 
Navigating to http://localhost:8000/pkg/testing/ would be worthwhile to see what's available to you.

If you don't have `godoc` command, then maybe you are using the newer version of Go (1.14 or later) which is no longer including `godoc`. 
You can manually install it with `go install golang.org/x/tools/cmd/godoc@latest`.


Other links

- https://blog.boot.dev/golang/best-ways-to-learn-golang/
- https://www.codewars.com/collections/golang-learned-katas


## Examples

An example can be added to a `_test.go` file and the output of the example will be added to the documentation. 
This allows you to document how a function behaves but the specification is also an executable test.

For example:

```go
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
```

Note the `// Output: 6` comment. This is used by the test function to check the output is correct and must be present to run the test. 


## Benchmarking

Benchmarks are little performance tests and can be useful to check if refactors have made your code faster or slower. 
They also go in a `_test.go` file. You can run benchmarks with `go test -bench=.`

For example:

```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
```

Note the `b *testing.B` argument provides access to `b.N` in the for loop. This is the number of times the benchmark function will run 
and is determined by the framework. The code in the for loop {} is what's being timed.

Sample output is:

```bash
❯ go test -bench=.
goos: linux
goarch: amd64
pkg: iteration
cpu: 11th Gen Intel(R) Core(TM) i7-11850H @ 2.50GHz
BenchmarkRepeat-16       6492337               178.2 ns/op
PASS
ok      iteration       1.352s
```

## Test Coverage

You can check test coverage with `go test -cover` this produces output like:

```bash
❯ go test -cover
PASS
        arrays  coverage: 100.0% of statements
ok      arrays  0.003s
```

## Interesting Language notes

### Arrays and slices

An interesting property of arrays is that the size is encoded in its type. If you try to pass an [4]int 
into a function that expects [5]int, it won't compile. They are different types so it's just the same as trying to pass a string into
a function that wants an int. You may be thinking it's quite cumbersome that arrays have a fixed length, and most of the time you 
probably won't be using them!

Go has slices which do not encode the size of the collection and instead can have any size.

You can take a slice of a slice with the pythonic style sytax `slice[low:high]` where low is the index of where to start the slice and 
high is the index where to end it (but not including the index itself). For example:

```go
numbers := []int{1, 2, 3, 4, 5}
numbers[1:3] // [2, 3]
numbers[:2] // [1, 2]
numbers[2:] // [3, 4, 5]
```

It is important to note that when slicing an array changing the slice affects the original array; 
but a "copy" of the slice will not affect the original array:

For example:

```go
package main

import (
  "fmt"
)

func main() {
  x := [3]string{"Bob", "Alice", "Jane"}

  y := x[:] // slice "y" points to the underlying array "x"

  z := make([]string, len(x))
  copy(z, x[:]) // slice "z" is a copy of the slice created from array "x"

  y[1] = "Suzy" // the value at index 1 is now "Belka" for both "y" and "x"

  fmt.Printf("%T %v\n", x, x)
  fmt.Printf("%T %v\n", y, y)
  fmt.Printf("%T %v\n", z, z)
```

when this code is run:

```text
[3]string [Bob Suzy Jane]
[]string [Bob Suzy Jane]
[]string [Bob Alice Jane]
```

Good blog post on slices https://blog.golang.org/go-slices-usage-and-internals

### First class functions

In go you can assign a function to a variable and pass it around like any other value.

This can be very useful to apply a functional model to your code to aid testability.

For example, you could write a function that takes a function as an argument and calls it for each item in a collection:

```go
someFunction := func(item int) {
    fmt.Println(item)
}
func Sum(numbers []int) int { 
  sum := 0
  for _, number := range numbers {
    someFunction(number)
    sum += number
  }
  return sum
}
```

### Interfaces

Interfaces are a way to define behaviour. They are a collection of method signatures that a type must implement in order to be considered an implementation of the interface.

Let's take the idea of a `Shape`, how does something become a shape? We just tell Go what a Shape is using an interface declaration:

```go
type Shape interface {
	Area() float64
}
```

Once you add this to the code, other types can now implement this interface by implementing the Area() method,
there is no requirement for a type to explicitly state that it implements an interface. In other words, in Go
interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.

### Pointers

Pointers are important in go to allow you to pass a reference to an object. 

For example take this definition of a wallet:

```go
type Wallet struct {
	balance int
}
func (w Wallet) Deposit(amount int) {
 	w.balance += amount
}
```

In Go, when you call a function or a method the arguments are copied.
So when calling `func (w Wallet) Deposit(amount int)` the `w`` is a copy of whatever we called the method from.
To fix this we can pass in a pointer to a wallet and operate on that object directly:

```go
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}
```

note the parameter is now `*Wallet` not `Wallet`. This means we are passing in a pointer to a wallet.

You may wonder why the pointer does not need to be dereferenced to create a `Wallet` object say with code like this:

```go
func (w *Wallet) Deposit(amount int) {
	(*w).balance += amount
}
```

Instead we seemingly addressed the object directly. In fact, the code above using `(*w)` is absolutely valid. However, the makers of Go deemed this notation cumbersome, so the language permits us to write w.balance, without an explicit dereference. These pointers to structs even have their own name: struct pointers and they are automatically dereferenced.

## Maps

Maps are a built in type in Go that are similar to dictionaries in Python. They are a collection of key value pairs.
Declaring a Map is somewhat similar to an array. Except, it starts with the map keyword and requires two types. 

The first is the key type, which is written inside the [].  it must be a type that can be compared with ==. See
https://golang.org/ref/spec#Comparison_operators for more details.

The second is the value type, which goes right after the []. This can be any type you like

For example:

```go
dictionary := map[string]int{"dave", 123}
```

You can add a value to a map as follows:

```go
dictionary["bob"] = 345
```

An interesting property of maps is that you can modify them without passing as an address to it (e.g &myMap)
This may make them feel like a "reference type", but [as Dave Cheney describes](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it) they are not.

> A map value is a pointer to a runtime.hmap structure.

So when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data.

A gotcha with maps is that they can be a nil value. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic. You can read more about maps [here](https://blog.golang.org/go-maps-in-action). Therefore, you should never initialize an empty map variable:

```go
var m map[string]string
```

Instead, you can initialize an empty map like we were doing above, or use the make keyword to create a map for you:

```go
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)
```

## Cannot import main: a Go Module gotcha

See article here: https://appliedgo.net/testmain/

TL;DR:

If you call your module main and have a test then run:

```bash
> go test
```

inside the project directory, but instead of the usual output, you get this:

```bash
# main.test
/var/folders/_m/dgnkqt8d3j10svk5c06px4vc0000gn/T/go-build306511963/b001/_testmain.go:13:2: cannot import "main"
FAIL    main [build failed]
```

Seems best to rename main package to something else.
