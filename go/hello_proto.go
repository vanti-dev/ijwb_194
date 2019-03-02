package main

import (
	"example.com/proto/helloworld"
	"fmt"
)

func main() {
	fmt.Print(helloworld.HelloReply{Message: "Hello, proto"})
}
