package agentHttp

import (
	"sync"

	"github.com/JieanYang/HelloWorldGoAgent/src/agentHttp/agentHttpController"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func StartHttp() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		router := gin.Default()

		router.GET("/", agentHttpController.HomeGetController)
		router.POST("/", agentHttpController.HomePostController)

		// Auth
		router.GET("/auth/authenticateByAuthKey", agentHttpController.HomeGetController)
		router.GET("/auth/generateTransferKeyByAuthKey", agentHttpController.HomeGetController)
		router.GET("/auth/generateSessionKeyByTransferKey", agentHttpController.HomeGetController)

		// RunCommand - with session key
		router.POST("/RunCommandByScriptContent", agentHttpController.RunCommandByScriptContent)
		router.POST("/RunCommandWithUrl", agentHttpController.RunCommandByUrl)

		// Run http web service
		endless.ListenAndServe(":9001", router)
		wg.Done()
	}()

	wg.Wait()

}
