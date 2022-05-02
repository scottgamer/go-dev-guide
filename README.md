# Go Dev Guide

## Commands

- `go build` - compiles a bunch of go source code files
- `go run` - compiles and executes one or two files
- `go fmt` - formats all the code in each file in the current directory
- `go install` - compiles and installs a package
- `go get` - downloads the raw source code of someone else's package
- `go test` - runs any tests associated with the current project

## Types of packages

- Executable - generates a file that we can run
- Reusable - code used as `helpers`, good place to put reusable logic

- The word `main` makes an executable type package
- When using the `go build` command, it checks if the package is `main` and generates an executable file

- In the other hand, if it's not `main`, it won't return anything
- The `main` package must have a function called `main` too

## Go standard libraries

- debug
- fmt
- io
- crypto
- encoding
- math

[Read more](https://pkg.go.dev/stdhttps://pkg.go.dev/std)

## Types and data structures

- Arrays
  - fixed size in length
- Slices
  - all items must be of same type
  - allows to grow or shrink in size

## Pointers

- &variable: returns the `memory address` of the `value` this variable is pointing at
- \*pointer: returns the `value` this `memory address` is point at
- turn `address` into `value` with `*address`
- turn `value` into `address` with `&value`
- **Note:** as long as we have the `*pointer` as the param in function, the `&address` is optional

```go
func main {
  var jim Person

  // option 1: getting the reference to the address
  jimPointer := &jim // gets memory address (reference) and assigns to jimPointer
	jimPointer.updateName("Jimmy")
	jim.print()

  // option 2: directly passing the value
  // this works as long as we have the * pointer as the param type
  jim.updateName("Jimmy")
	jim.print()

}

func (pointerToPerson *Person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
	// pointerToPerson.firstName = newName // this also works
}
```

### Pointers and Slices

- `Slices` behave different to any other `primitive` or `struct`
- the reason for this is that internally Go allocates 2 directions in memory
  - one for the slice `metadata` (length, capacity, ptr to head)
  - the other for the `array` of elements
- when passing the slice to a function, it makes a copy of the `metadata` (as a pass by value)
- but, since the `metadata's ptr to head` points to the same address the `array` is, the change is made by reference
- in this case, there's no need to pass a `pointer` to the function

### Value types & Reference types

| Value Types | Reference Types |
| ----------- | --------------- |
| int         | slices          |
| float       | maps            |
| string      | channels        |
| bools       | pointers        |
| structs     | functions       |
