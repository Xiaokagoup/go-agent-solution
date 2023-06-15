//
// Copyright (C) 2023 ANSYS, Inc. Unauthorized use, distribution, or duplication is prohibited.
//

package main

import (
	"fmt"

	agt "AnsysCSPAgent/src/agent"
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

	agent := agt.NewAgent()

	agent.Launch()

	fmt.Println("Program main func - end")

	// os.Exit(0) // Close agent
}
