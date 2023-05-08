package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	agt "github.com/JieanYang/HelloWorldGoAgent/src/agent"
	docs "github.com/JieanYang/HelloWorldGoAgent/src/docs"
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

// @host      localhost:9001
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	host := os.Getenv("HOST")
	fmt.Println("host", host)

	docs.SwaggerInfo.Host = host

	fmt.Println("Hello World !")
	agent := agt.NewAgent()

	agent.Init() // Init agent
	agent.Start()

	// os.Exit(0) // Close agent
}
