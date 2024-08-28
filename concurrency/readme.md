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