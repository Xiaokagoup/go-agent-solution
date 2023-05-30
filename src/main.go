package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	agt "AnsysCSPAgent/src/agent"
	docs "AnsysCSPAgent/src/docs"
)

// @title           AnsysCSPAgent API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @host      localhost:9001
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	fmt.Println("Program main func - start")

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	host := os.Getenv("HOST")
	fmt.Println("host", host)

	docs.SwaggerInfo.Host = host

	fmt.Println("Hello World !")
	agent := agt.NewAgent()

	agent.Launch()

	fmt.Println("Program main func - end")

	// os.Exit(0) // Close agent
}
