### Concurrency
__When a Go program starts, it is allocated Logical Processors (P), which are software constructs managed by Go's runtime to execute Go routines. Here's a breakdown of the explanation:__

**Logical Processor (P):** This is a virtual processor created by the Go runtime to manage concurrency. Each P can run one Go routine at a time.

**Virtual Core:** In this context, a virtual core refers to how the operating system presents CPU resources to programs. Each physical CPU core might have multiple hardware threads (as in Hyper-Threading, where one physical core can execute multiple threads in parallel). The operating system treats each of these threads as a separate virtual core.

**Multiple Hardware Threads (Hyper-Threading):** If your processor supports Hyper-Threading, it means each physical core can run more than one thread at a time (often two). These hardware threads are treated by Go as virtual cores.

**Go Program and Virtual Cores:** Go's runtime assigns a logical processor (P) to each virtual core. For example, if your system has 4 physical cores with Hyper-Threading enabled (2 threads per core), your Go program will see and use 8 virtual cores.






#### Normal Running a function
```go
statment1
start()
statement2
```
In the normal running of a function for the above scenario.

First, statement1 will be executed
Then start() function will be called
Once the start() function finishes then statement2 will be executed

#### Running a function as a goroutine
```go
statment1
go start()
statement2
```
In running a function as a goroutine for the above scenario

First, statement1 will be executed
Then function start() will be called as a goroutine which will execute asynchronously.
statement2 will be executed immediately. It will not wait for start() function to complete. The start function will be executed concurrently as a goroutine while the rest
of the program continues its execution.

So basically when calling a function as a goroutine, call will return
immediately the execution will continue from the next line while the goroutine will be executed concurrently in the background. **_Also note that any return value from the goroutine will be ignored._**

Goroutines don’t have parents or children.

The scheduler also gets the opportunity for contexts switch on below events too

- Functions Call
- Garbage Collection
- Network Calls
- Channel operations
- On using go keyword
- Blocking on primitives such as mutex etc

this way deadlocks and race conditons are avoided. Infact go believes in the mantra
"Don't share memory
for communication, instead share memory by communicating"



```go
package main

import (
	"fmt"
	"runtime"
)

func main() {

    // NumCPU returns the number of logical
    // CPUs usable by the current process.
    fmt.Println(runtime.NumCPU())
}
```



