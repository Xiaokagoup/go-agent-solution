package agentHttp

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type RequestData struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

func StartHttp() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		router := gin.Default()

		router.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		// RunCommand
		router.POST("/", func(c *gin.Context) {
			var reqData RequestData
			if err := c.ShouldBindJSON(&reqData); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"results": reqData})
		})

		// Run http web service
		router.Run(":9001")
		wg.Done()

		// http.HandleFunc("/", agentHttpController.HomeController)
		// // http.HandleFunc("/RunCommandWithFormData", agentHttpController.RunCommandWithFormData)
		// // http.HandleFunc("/RunCommandWithBody", agentHttpController.RunCommandWithBody)

		// // Auth
		// http.HandleFunc("/auth/authenticateByAuthKey", agentHttpController.HomeController)
		// http.HandleFunc("/auth/generateTransferKeyByAuthKey", agentHttpController.HomeController)
		// http.HandleFunc("/auth/generateSessionKeyByTransferKey", agentHttpController.HomeController)

		// // RunCommand - with session key
		// http.HandleFunc("/RunCommandByScriptContent", agentHttpController.RunCommandByScriptContent)
		// http.HandleFunc("/RunCommandWithUrl", agentHttpController.RunCommandByUrl)

		// err := http.ListenAndServe(":9001", nil) // Block code
		// if err != nil {
		// 	panic(err)
		// }
	}()

	wg.Wait()

}
