package agentHttp

import (
	"sync"

	"github.com/JieanYang/HelloWorldGoAgent/src/agentHttp/agentHttpController"
	// "github.com/fvbock/endless"
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
		router.Run(":9001")
		wg.Done()
	}()

	wg.Wait()

}
