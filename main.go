package main

import (
	"fmt"

	appsync "github.com/sony/appsync-client-go"
)

func main() {
	fmt.Println("Hello, playground")
	appsync.Subscriber()
}
