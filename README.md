# globocop
A static analysis tool that checks for global variables in Go code.

## Installation


## Usage

globocop comes with 2 example packages:

* ./examples/with_global - a package that contains global identifiers.
```
$ go run globocop.go ./examples
/.../examples/example.go:6:2: global var "unexportedVariable"
/.../examples/example.go:10:2: global var "ExportedVariable"
exit status 3
```