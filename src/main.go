package main

import (
	"fmt"

	agt "github.com/JieanYang/HelloWorldGoAgent/src/agent"
	_ "github.com/JieanYang/HelloWorldGoAgent/src/docs"
)

func main() {

	fmt.Println("Hello World !")
	agent := agt.NewAgent()

	agent.Start()

	// os.Exit(0) // Close agent
}
