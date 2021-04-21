package main

import (
	"fmt"
	//"./src/github.com/dreinix/gonake.backend/pkg/listing/"
	Router "github.com/dreinix/gonake/pkg/router"
)

func main() {
	fmt.Println("Start server")
	Router.StartServer()
}
