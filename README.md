# learngowithtests

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