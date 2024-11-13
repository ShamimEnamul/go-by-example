	# declarations
	var str1 string = "first string"
	str2 := "second string"
	fmt.Println(str1, " : ", str2)


### string is immutable

```go
mystr := "Welcome to GeeksforGeeks"
fmt.Println("String:", mystr)

/* if you trying to change
      the value of the string
      then the compiler will
      throw an error, i.e,
     cannot assign to mystr[1]

   mystr[1]= 'G'
   fmt.Println("String:", mystr)
*/
```


### Interate through stirng

```go
for indx, chr := range "helloworld" {
//fmt.Println(indx, chr)
fmt.Printf("%c and %d\n", chr, indx)
}
```


```go
for indx, chr := range "helloworld" {
fmt.Printf("%c and %d\n", chr, indx)
}
```
chr represents each character in the string "helloworld". It is of type rune, which is an alias for int32 and represents a Unicode code point.
indx represents the index of the character in the string.
fmt.Printf("%c and %d\n", chr, indx) does the following:

%c converts the rune value (which is an integer representing a Unicode code point) into its corresponding character and prints it.
%d prints the index (indx) as a decimal number.
So, the code outputs each character of "helloworld" alongside its index in the string.

fmt.Println prints the raw chr value, which is the Unicode code point (in integer form) of the character, rather than the character itself.

If you replace fmt.Printf with fmt.Println(indx, chr), it will print the index and the corresponding Unicode code point of each character:

```go
for indx, chr := range "helloworld" {
fmt.Println(indx, chr)
}
```

For example:
The character 'h' has a Unicode code point of 104, so it would print 0 104 for the first iteration.

### what is unicode point?

A Unicode code point is a unique number assigned to each character in the Unicode standard, which is a global encoding system designed to represent text from all writing systems, symbols, and even emojis.

The code point is usually written in the format U+ followed by a hexadecimal number. For example:

The letter 'A' has a Unicode code point of U+0041.
The character '你' (a Chinese character) has a Unicode code point of U+4F60.
The emoji '😊' has a Unicode code point of U+1F60A.
In Go, a rune represents a Unicode code point, allowing you to work with any character, regardless of the language or symbol set it belongs to.