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