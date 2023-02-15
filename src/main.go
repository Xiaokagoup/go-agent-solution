package main

import (
	"fmt"

	// agt "github.com/JieanYang/HelloWorldGoAgent/src/agent"
	agtHttp "github.com/JieanYang/HelloWorldGoAgent/src/agentHttp"
)

func main() {

	fmt.Println("Hello World !")
	// agent := agt.NewAgent()
	agtHttp.StartHttp()
	// agent.Start()

}
