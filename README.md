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
Go supports function clojure superhappyfuntime!
1. Func Expression : You can declare function in another function scope as "funcname := func() int {func logic}".
2. Normal Func : You can declare function outside main as "func (x int) funcname() int {func logic}"
3. Return a Func : you can return a func when declared "func wrapper() func() int {func logic}"
