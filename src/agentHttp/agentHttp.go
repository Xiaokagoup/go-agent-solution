package agentHttp

import (
	"sync"

	"AnsysCSPAgent/src/agentHttp/agentHttpController"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartHttp() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		router := gin.Default()
		config := cors.DefaultConfig()
		config.AllowAllOrigins = true
		router.Use(cors.New(config))

		// Swagger
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		router.GET("/", agentHttpController.HomeGetController)
		router.POST("/", agentHttpController.HomePostController)

		// Auth
		router.GET("/auth/authenticateByAuthKey", agentHttpController.HomeGetController)
		router.GET("/auth/generateTransferKeyByAuthKey", agentHttpController.HomeGetController)
		router.GET("/auth/generateSessionKeyByTransferKey", agentHttpController.HomeGetController)

		// RunCommand - with session key
		router.POST("/RunCommandByScriptContent", agentHttpController.RunCommandByScriptContent)
		router.POST("/RunCommandWithUrl", agentHttpController.RunCommandWithUrl)

		// Exits agent
		router.GET("/Exit", agentHttpController.Exit)

		// Get AppConfig
		router.GET("/dev/getAppConfig", agentHttpController.GetAppConfig)

		// Test
		router.GET("/dev/test", agentHttpController.Test)

		// Run http web service
		router.Run(":9001")
		wg.Done()
	}()

	wg.Wait()

}
