package agentHttp

import (
	"sync"

	"github.com/JieanYang/HelloWorldGoAgent/src/agentHttp/agentHttpController"
	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
func StartHttp() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		router := gin.Default()
		config := cors.DefaultConfig()
		config.AllowAllOrigins = true
		router.Use(cors.New(config))

		router.GET("/", agentHttpController.HomeGetController)
		router.POST("/", agentHttpController.HomePostController)

		// Auth
		router.GET("/auth/authenticateByAuthKey", agentHttpController.HomeGetController)
		router.GET("/auth/generateTransferKeyByAuthKey", agentHttpController.HomeGetController)
		router.GET("/auth/generateSessionKeyByTransferKey", agentHttpController.HomeGetController)

		// RunCommand - with session key
		router.POST("/RunCommandByScriptContent", agentHttpController.RunCommandByScriptContent)
		router.POST("/RunCommandWithUrl", agentHttpController.RunCommandWithUrl)

		// Swagger
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		// Run http web service
		endless.ListenAndServe(":9001", router)
		wg.Done()
	}()

	wg.Wait()

}
