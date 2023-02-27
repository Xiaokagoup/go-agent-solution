package main

import (
	"fmt"

	agt "github.com/JieanYang/HelloWorldGoAgent/src/agent"
	_ "github.com/JieanYang/HelloWorldGoAgent/src/docs"
)

// @title           HellowWorldGoAgent API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      *:9001
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	fmt.Println("Hello World !")
	agent := agt.NewAgent()

	agent.Start()

	// os.Exit(0) // Close agent
}
