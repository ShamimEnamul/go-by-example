### buffer 

```prototext
syntax = "proto3";

package main;
option go_package = "./";

message SocialMedia{
  string youtube = 1;
  string twitter = 2;
}
message Person{
  string name = 1;
  int32 age = 2;
  SocialMedia  socialMedia = 3;
}
```

```go
package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"log"
)

func main() {
	fmt.Printf(">> Buffer")
	shamim := &Person{
		Name: "shamim",
		Age:  24,
		SocialMedia: &SocialMedia{
			Youtube: "youtube",
			Twitter: "twitter",
		},
	}

	data, err := proto.Marshal(shamim)
	if err != nil {
		log.Fatal("Marshaling error", err)
	}

	fmt.Println(data)
	data2 := &Person{}
	err = proto.Unmarshal(data, data2)
	if err != nil {
		log.Fatal("Unmarshalling error occurred", err)
	}

	fmt.Println(data2.GetName())
	fmt.Println(data2.GetAge())
	fmt.Println(data2.GetSocialMedia())

}

```

This code is an example of how to use Protocol Buffers (protobuf) in a Go program. Protocol Buffers are a method for serializing structured data, similar to XML or JSON, but more efficient and faster. In this code, a simple Person message with associated SocialMedia is defined in a .proto file, and Go code is generated to work with this message format.

Here's a step-by-step explanation:

Importing Libraries: The code imports necessary libraries for using Protocol Buffers ("google.golang.org/protobuf/proto") and standard Go packages ("fmt" and "log").

Message Definition (person.proto): The .proto file defines two messages, SocialMedia and Person. The Person message has fields for name (a string), age (an int32), and a SocialMedia object. These messages are defined using Protocol Buffers syntax, and the go_package option specifies the Go package name where the generated code will be placed.

main() Function:

A Person object named shamim is created with sample data.
proto.Marshal is used to serialize the shamim object into binary data stored in the data variable.
fmt.Println(data) prints the serialized binary data to the console.
A new Person object, data2, is created.
proto.Unmarshal is used to deserialize the data into the data2 object.
fmt.Println is used to print the individual fields of the deserialized data2 object: name, age, and SocialMedia.
Output:

The serialized binary data (usually not human-readable) is printed first.
The deserialized data is then printed, which should match the original shamim data.
This example demonstrates the basic usage of Protocol Buffers in Go for serializing and deserializing structured data. Protocol Buffers are especially useful for defining structured data formats and exchanging data between different systems while minimizing the size of the data transferred. The .proto file is used to define the message structure, and Go code generated from the .proto file provides methods to work with these messages in your Go programs.