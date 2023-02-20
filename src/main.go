package main

import (
	"fmt"

	agt "github.com/JieanYang/HelloWorldGoAgent/src/agent"
)

func main() {

	fmt.Println("Hello World !")
	agent := agt.NewAgent()

	agent.Start()

}
