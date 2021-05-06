package main

import (
	"fmt"

	Router "github.com/dreinix/gonake/pkg/router"
)

func main() {
	fmt.Println("Start server")
	Router.StartServer()
}
