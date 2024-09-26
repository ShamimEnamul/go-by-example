# MEMORY MANAGEMENT

### When are values constructed on the heap
1. When a value could possibly be referenced after the function that constructed the value returns
2. When the compiler determines a value is too large to fit on the stack
3. When the compiler doesn't know the size of the variable such as slice

### Some commonly allocated variables
1. Values are share with pointers
2. Variables stored in interface variables
3. Func literal variables
   - Variables captured by a closure
4. Backing data for Maps, Channels, Slices and Strings

### REFERENCES
- https://www.youtube.com/watch?v=ZMZpH4yT7M0
