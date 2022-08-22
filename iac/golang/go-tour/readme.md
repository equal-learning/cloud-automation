# A tour of Go

## Package variable constant functions
- package
- import
- export
- basic (primitive/built-in) types
- variables
- variable declaration shorthand  
- variable with initializers
- constants
- uninitialized variables
- type inference
- type conversion
- numerical constants
- functions
- functions with named return values
- functions with multiple return values
- function values
- function closure

## Flow control : if else for defer switch
- if else
- if with a short statement
- for
- for ever : for {}
- for as while
- defer executing to after the return of the function call
- defer stacking (lifo)
- switch
- switch with no condition

## pointer struct slice map
- pointer
- array
- slice
- slice length & capacity
- slice nil
- slice range
- slice pointer
- slice making
- slice literals
- slice append
- map
- map mutation
- map literal
- struct composite type
- struct fields
- struct literals
- struct tags
- struct pointers

## method
- method : a function with struct type receiver
- method : a function with primitive (non-struct) type receiver
- method with value receiver
- method with pointer receiver
- No class in go
- interface type (value,type) : Holds any type that implements the methods of the interface
- interface empty : a generic type holder
- interface with nil value : throws an error if called
- interface single type assertion : i.(T)
- interface several type assertion : type switch
- interface Stringer built-in : equivalent to toString
- interface Error : to handle error

## Concurrency
- routine : thread in go (runs in a shared address space)
- routine mutext : lock unlock to implement mutual exclusivity
- channel chan <type> :  typed conduite through which values can be sent and recieve synchronously
- channel with buffer
- channel closing
- routine select : wait on multiple coditions
- routine select default : run when no waiting condition is ready


Reference ::  

[Coding Guideline](https://go.dev/doc/effective_go)  
[Testing Guideline](https://go.dev/doc/code)  
[Packaging Guideline](https://pkg.go.dev/cmd/go)  
