# Learn Go

https://github.com/GoesToEleven/GolangTraining
https://github.com/ardanlabs/gotraining

### Go paths need to be defined before anything

* export GOPATH=$HOME/work/
* export GOROOT=/usr/local/go
* export PATH="$HOME/Documents/work/bin:$PATH"

### Use atom as a text editor

1. get atom : https://atom.io/
2. get go-plus plugin : https://github.com/joefitzgerald/go-plus
* just open atom and hit ctrl-shit-p search for packages to install and get go-plus
* now you have code completion

### fmt

1. Formating library, pronounced, fmt - not ef-em-tee
2. fmt.Println() - variatic function, its smart, you can print different types in one statement without casting
3. fmt.Scan : needs to point to memory address (not the var). I store my name " fmt.Scan(&myName) "

### go commands

1. "go build" : build and executable file run it with  ./filename
2. "go clean" : removes the executable files
3. "go install" : make program an environment variable, then you can call it anywhere by just writing the filename. places a executable program in you environment bin folder and .a file in your archives
4. "go get" :  get new packages [-d] [-f] [-fix] [-insecure] [-t] [-u] package-name
5. "go run" : looks for a main run executable

### Encoding https://godoc.org/fmt
1. Binary : fmt.Printf("%b")
2. Hexadecimal : fmt.Printf("%x")
3. Decimal : fmt.Printf("%d")
4. Unicode : fmt.Printf("%q")

### Package conventions

Be careful of package idioms! Conventions are strongly enforced

1. Comment : "Package packagename" before package declaration
2. Comment : "Functionname" before function declaration -- same for package variables
3. Capitalized functions/variables from packages!
4. Fully qualified name needed for 3rd party libs

### Variables

Be idiomatic with go variables. Don't be some kind of space cowboy.

1. Shorthand notation: a := 1 --- b := "hello world"
2. Initializes to 0 : var a int --- var b float64
3. Only use "var a int8 = 4" if you don't want default size
4. Type format "%T"
5. Multiple initialization "var a,b int 8 = 34,54"

### Scope

Mostly works like all other languages. but try to keep you scope tight. don't unnecessarily give large scope! obviously...
In order of restrictiveness.

1. Universe : good for the whole shebang
2. Package : scope, withing package called
3. file : a top of file. this does not support shorthand assignment of variables
4. Block : within block only

### Functions

Go supports function clojure superhappyfuntime! Everything in on is pass by value!!

1. Func Expression : You can declare function in another function scope as "funcname := func() int {func logic}".
2. Normal Func : You can declare function outside main as "func (x int) funcname() int {func logic}"
3. Return a Func : you can return a func when declared "func wrapper() func() int {func logic}"
4. Callbacks : function that takes a function as a parameter - similar to a lambda expression (functional methodology - haskell,lisp)

### Http package - "net/http"
easily do http requests. can easily be go'ed (threaded)

### Constant
Golang supports constant - functional func func func.
Can be a kind (which means it just picks the ideal type) or a type that is explicitly defined. this creates some measure of freedom. and slightly less tedious to code.

1. "const val = 42"
2. "const ( pi = 3.14 languague = "go")"
3. iota is a automatic incrementer. kind of like an ID. only takes 8 bits of space!
4. check with import "unsafe" --- unsafe.Sizeof(const/var)

### Memory Addressing
go support memory Addressing

1. var a = 12 - is at address &a - &a is a pointer to an address - \*&a is the content of that address
2. change values with pointers a := 1 -- var b = &a -- \*b = 3. Now a = 3. This means you only have to more around pointers instead of moving actually data. makes things much faster! "pass by value!"
3. variable a has address &a can be pointed to \*a and must be passed into function myfunc(a \*int)

### Statements
1. Defer : runs command at the end of main function, quite useful for keeping code clean, you can open a file and defer its close right under it.
2. make : creates slices, maps, channels only. Make creates a reference to a data structure, so when you pass them into a function (actually passing the regerence). they can modify an existing underlying DS

### Data Structures
1. Array : numbered sequence - not idiomatic || var myarray [10]int
2. slice : list descriptor of contiguous segment - initializes to nil, changeable size and its length and capacity can be different || make([]int, 5, 10). Supports append, slice is just a dynamic array. when initialized with no values, the slice is nil, because there is no underlying data structure with attributes (pointer, len, cap).
3. map : map... dict... wtv you wanna call it || make(map[string]int). map is also a reference type! can be created with composite litteral syntax map[int]int{}.
4. struct : sequence of fields with name and type. Composite DS type (like a class, but don't say that too loud) || type potato struct{}. You can also have unnamed fields. Encapsulation (exported/unexported), Re-usability (inheritance), Polymorphism (interfaces), Overriding(promotion). Instead of instantiating we say creating a value of a type! when declaring a struct you can add tags to some fields (for example if you don't want them to be stored)
