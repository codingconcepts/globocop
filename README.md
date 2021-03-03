# globocop
A static analysis tool that checks for global variables in Go code.

## Installation


## Usage

globocop comes with 2 example packages:

* ./examples/with_global - a package that contains global identifiers.
```
$ go run globocop.go ./examples/with_global
/.../examples/with_global/example.go:5:2: global variable or constant unexportedVariable
/.../examples/with_global/example.go:8:2: global variable or constant ExportedVariable
/.../examples/with_global/example.go:13:2: global variable or constant unexportedConstant
/.../examples/with_global/example.go:16:2: global variable or constant ExportedConstant
exit status 3
```

* ./examples/without_global - a package that doesn't contain global identifiers.
```
$ go run globocop.go ./examples/without_global
exit status 0
```