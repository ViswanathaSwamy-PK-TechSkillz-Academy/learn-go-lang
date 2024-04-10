# Learn `Go` Language

I am learning `Go` Language from different Video Courses, Books, and Websites. We will be learning

## Topics Planned

### Session 1

> 1. The Big Picture
>    - Previous Session(s)
> 1. Series Introduction
> 1. What is `Go`?
> 1. The `Go` Playground
> 1. `Go` installation
>    - Verifying `Go` version
>    - Retrieving available `Go` commands
>    - Getting Help on `Go` Commands
> 1. Install `VS Code`, and `golang.go` VS Code extension
> 1. Install/Update `Go tools` in VS Code
> 1. Discussion on a couple of `Go tools`
>    - `gopls`
>    - `staticcheck`
>    - `goplay`
> 1. Few `Go` commands
>    - `go mod init`
>    - `go run .`
> 1. Let's `Go` with our first `Go` program
>    - Create a `Go` module inside `helloworld` folder
>    - Create a main.go file, to contain the code
>    - Quick Code Walk-through
>    - Executing the `Go` code i.e., main.go file
>    - Create a go.mod and main.go file, inside another folder `helloworldv1`
>    - Executing the `Go` code
>    - Few ways to execute `Go` code
> 1. SUMMARY / RECAP / Q&A

### Session 2

> 1. The Big Picture
>    - Previous Session(s)
> 1. `Go` tool chain
> 1. Single Binary
>    - `.exe` location when executing from `VS Code`
>    - `.exe` location when executing using `go run`
>    - `.exe` location when executing using `go build`
> 1. Values Demo
> 1. Escape sequences
> 1. Basic Types
> 1. Variables
> 1. Type Conversions
> 1. Arithmetic
> 1. Comparison Operators
> 1. Constants, Constant Expression
> 1. SUMMARY / RECAP / Q&A

### Session 3

> 1. Functions - Take no arguments, and returns nothing
> 1. Basic Types Part II (Complex numbers, rune, character)
> 1. Constants, Constant Expression, Iota
> 1. Pointers
> 1. Functions - Take arguments, and returns nothing
> 1. Functions - Take arguments, same type can ignore mentioning for each variables

### Session 4

> 1. Functions - Take no arguments, and returns value
> 1. Functions - Take no arguments, and returns multiple values
> 1. Functions - Take arguments, and returns value

### Session 5

> 1. Debugging

```go
// Character
a, b := 'A', '世'
fmt.Println(a, string(a))
fmt.Println(b, string(b))
fmt.Printf("%c %c\n", a, b)
```

---

## A Tour of Go

|| Name 1    | Name 2 | Name 3    | Name 4 | Name 5   |
|-------- | -------- | ------- | -------- | ------- | -------- |
|1. | Packages | Imports | Exported names | Functions | Functions continued |
|2. | Multiple Results | Named return values | Variables | Variables with initializers | Short variable declarations |
|3. | Basic types | Zero values | Type conversions | Type inference | Constants |
|4. | Numeric Constants | For | For continued | For is Go's "while" | Forever |
|5. | If | If with a short statement | If and else | Exercise: Loops and Functions | Switch |
|6. | Switch evaluation order | Switch with no condition | Defer | Stacking defers | Pointers |
|7. | Structs | Struct Fields | Pointers to structs | Struct Literals | Arrays |
|8. | Slices | Slices ... references to arrays | Slice literals | Slice defaults | Slice length and capacity |
|9. | Nil slices | Creating a slice with make | Slices of slices | Appending to a slice | Range |
|10. | Range continued | Exercise: Slices | Maps | Map literals | Map literals continued |
|11. | Mutating Maps | Exercise: Maps | Function values | Function closures | Exercise: Fibonacci closure |
|12. | Methods | Methods are functions | Methods continued | Pointer receivers | Pointers and functions |
|13. | Methods and pointer indirection | Methods and pointer indirection (2) | Choosing a value or pointer receiver | Interfaces | Interfaces are implemented implicitly |
|14. | Interface values | Interface values with nil underlying values | 3 | 4 | 5 |
|15. | 1 | 2 | 3 | 4 | 5 |

## Go by Example

|| Name 1    | Name 2 | Name 3    | Name 4 | Name 5   |
|-------- | -------- | ------- | -------- | ------- | -------- |
|1. | Hello World | Values | Variables | Constants | For |
|2. | If/Else | Switch | Arrays | Slices | Maps |
|3. | Range | Functions | Multiple Return Values | Variadic Functions | Closures |
|4. | Recursion | Pointers | Strings and Runes | Structs | Methods |
|5. | Interfaces | Struct Embedding | Generics | Errors | Goroutines |
|6. | Channels | Channel Buffering | Channel Synchronization | Channel Directions | Select |
|7. | Timeouts | Non-Blocking Channel Operations | Closing Channels | Range over Channels | Timers |
|8. | Tickers | Worker Pools | WaitGroups | Rate Limiting | Atomic Counters |
|9. | Mutexes | Stateful Goroutines | Sorting | Sorting by Functions | Panic |
|10. | Defer | Recover | String Functions | String Formatting | Text Templates |
|11. | Regular Expressions | JSON | XML | Time | Epoch |
|12. | Time Formatting / Parsing | Random Numbers | Number Parsing | URL Parsing | SHA256 Hashes |
|13. | Base64 Encoding | Reading Files | Writing Files | Line Filters | File Paths |
|14. | Directories | Temporary Files and Directories | 3 | 4 | 5 |
|15. | 1 | 2 | 3 | 4 | 5 |
