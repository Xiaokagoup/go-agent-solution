package agentHttpController

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"AnsysCSPAgent/src/tools/4_base/TRequest"
	"AnsysCSPAgent/src/tools/TRunCommand"
	"AnsysCSPAgent/src/tools/agentMetadataManager"

	"github.com/gin-gonic/gin"
)

func HomeGetController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong yang",
	})
}
func HomePostController(c *gin.Context) {
	var reqData TRunCommand.RunCommandWithUrlRequestData
	if err := c.ShouldBindJSON(&reqData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// fmt.Println("reqData", reqData)

	c.JSON(http.StatusOK, gin.H{"results": reqData})
}

// @Summary Run command using script content
// @Description description
// @Accept  json
// @Produce  json
// @Param object body RequestDataForRunCommandByScriptContent true "param for RunCommandByScriptContent"
// @Success 201 {string} string "The object was created successfully"
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Failed to create object"
// @Router /RunCommandByScriptContent [post]
func RunCommandByScriptContent(c *gin.Context) {
	var reqData TRunCommand.RequestDataForRunCommandByScriptContent
	if err := c.ShouldBindJSON(&reqData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("reqData", reqData)
	scriptOutput, _ := TRunCommand.RunCommandByScriptContent(string(reqData.ScriptContent))

	data := TRequest.StringRequestData{Result: string(scriptOutput)}

	c.JSON(http.StatusOK, gin.H{"results": data})

}

// @Summary Run command using url
// @Description description
// @Accept  json
// @Produce  json
// @Param object body RequestDataForRunCommandByUrl true "param for RunCommandWithUrl"
// @Success 201 {string} string "The object was created successfully"
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Failed to create object"
// @Router /RunCommandWithUrl [post]
func RunCommandWithUrl(c *gin.Context) {
	var reqData TRunCommand.RunCommandWithUrlRequestData
	if err := c.ShouldBindJSON(&reqData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("reqData", reqData)
	fmt.Println("reqData.Url", reqData.Url)

	data, err := TRunCommand.RunCommandByUrl(reqData.Url)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": data})

}

// @Summary Exit the agent
// @Description Exit the agent
// @Accept  json
// @Produce  json
// @Success 201 {string} string "The object was created successfully"
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Failed to create object"
// @Router /Exit [get]
func Exit(c *gin.Context) {
	fmt.Println("Exit API called")

	time.AfterFunc(1*time.Second, func() {
		os.Exit(0) // without error
		// os.Exit(1) // with error
	})

	c.JSON(http.StatusOK, gin.H{"results": "ok"})

	fmt.Println("End API called")
}

// @Summary Get origin metadata json
// @Description Get origin metadata json
// @Accept  json
// @Produce  json
// @Success 201 {string} string "The object was created successfully"
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Failed to create object"
// @Router /dev/getAppConfig [get]
func GetAppConfig(c *gin.Context) {
	config := agentMetadataManager.GetOrCreateConfigFile()
	c.JSON(http.StatusOK, gin.H{"results": config})
}

// @Summary Test
// @Description Test
// @Accept  json
// @Produce  json
// @Router /dev/test [get]
func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"results": "OK"})
}
