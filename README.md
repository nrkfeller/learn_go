# Learn Go
** Don't communicate by sharing memory; instead, share memory by comunicating **

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
5. If you want to export types or their attricutes, capitalize what is to be exported
6. func init(){}. runs before anything on that file or program

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

### Ecoding/Decoding -- Un/Marshalling
1. Encoding/Decoding : writing and reading into a stream. Typically used for external interaction.  Use json.NewEncoder(io.writer), gives out a pointer to a new \*Encoder. json.NewDecoder(io.reader), gives out pointer to \*Decoder
2. Marshalling and UnMarshalling : writing and reading from withing an application. use ticks /`/` to create a reader from a string and Unmarshal with the json.Unmarshal command. json.Marshal(type) to marshal type into a []byte

### Interfaces - substitutability/Polymorphism
1. Interface are abstract types. Functions should match on structs that want to use the same interface.
2. Polymorphism allows code to have different behavior through the implementation of types. Interfaces are types that declare behavior, the behavior is however not implemented/used by the interface but by the types that use that interface.
3.  When interface methods are called. the equivalent method from that specific type is instead evoked. This is polymorphism in action.
4. Implementation are satisfied implicitly! as long as the types has identical called and returned methods.
5. Remember concrete types have typed and values.
6. Sorting from the standard lib works through interfaces that can be over written, it all happens in place through the references.
7. Empty interface : all types implement the empty interface, you can place anything into a slice on empty interface but this may cause problems when unpacking different types that cannot all be treated the same
8. Assertion : You cannot cast interfaces, you need to assert the instead. type potato interface{} = "foobar" -- potato.(string)

### Concurrency
1. WaitGroup : var wg sync.WaitGroup. Acts like a semaphore. essentially available resources can be incremented and decremented!
2. GOMAXPROCS : above go version 1.5. Your go routines default to using all the core. You no longer need to specify it.
3. Mutex : var m sync.Mutex. m.Lock() m.Unlock(). Behavior is as expect. Like in all other languages
4. Atomicity : atomic.AddInt64(&counter, 1). Only one process can access this single value at the time. Like in all other languages.
5. Channels : channels are a way to pass information. channels can have information passed to them and removed from them using this operator <-. Passing information in a channel does not prevent functions or programs from exiting, this must be done explicitly.
* close(type channel) : prevents the program from putting new information into the channel. but can still receive. We need to close otherwise, we cannot iterate through all available data in the channel (ie deadlock will be created).
* ranging : ranging through a channel allows us to grab all the info off it without the program exiting. however it require the channel to be closed

### Channels
* Channels are a conduit which you can send and receive values with. c := make(chan int) || can be of any type.
* <- arrow is optional and shows the direction of the flow of data (bi-directional, send-only, receive-only). Bi-directional if omitted.

### Pipelining
A series of stages connected by channels. Each stage is a group of goroutines.
1. receive values from upstream via an inbound channel. essentially a bunch of instructions and values
2. Perform the work, compute the instructions on the values
3. Send the work downstream via outbound channels
use gen() to create the data and instructions. the a second descriptive function that distributes the work.

### Fanning in/out
Fan out is the process by which work is distributed to slaves to do the computation. Fan in it the process by which the work done by slaves is being consolidates back into a single channel.
