# globocop
A static analysis tool that checks for global variables in Go code.

## Installation

```
$ go get -u github.com/codingconcepts/globocop
```

## Usage

```
$ globocop ./examples
/.../examples/example.go:6:2: global var "unexportedVariable"
/.../examples/example.go:10:2: global var "ExportedVariable"
exit status 3
```