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
