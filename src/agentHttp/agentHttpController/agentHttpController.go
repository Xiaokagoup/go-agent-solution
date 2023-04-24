package agentHttpController

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/JieanYang/HelloWorldGoAgent/src/tools/agentOriginMetadataJsonManager"
	"github.com/JieanYang/HelloWorldGoAgent/src/tools/runCommand"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Reponse struct {
	Results string
}

type RequestDataForRunCommandByScriptContent struct {
	ScriptContent string `json:"scriptContent" default:"#!/bin/bash\necho \"start\"\necho \"hello yang\"\necho \"end\""`
	// ScriptContent string `json:"scriptContent" default:"Write-Output \"Windos PowerShell\"\nWrite-Output \"start\"\nWrite-Output \"hello yang\"\nWrite-Output \"end\""`
}
type RequestDataForRunCommandByUrl struct {
	Url string `json:"url" default:"https://ansys-gateway-development.s3.eu-west-3.amazonaws.com/2023-02-27-linux-script.sh"`
	// Url string `json:"url" default:"https://ansys-gateway-development.s3.eu-west-3.amazonaws.com/2023-02-27-windows-script.ps1"`
}

type RequestData struct {
	Value string `json:"value"`
	RequestDataForRunCommandByUrl
	RequestDataForRunCommandByScriptContent
}

func HomeGetController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong yang",
	})
}
func HomePostController(c *gin.Context) {
	var reqData RequestData
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
	var reqData RequestDataForRunCommandByScriptContent
	if err := c.ShouldBindJSON(&reqData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("reqData", reqData)
	scriptOutput := runCommand.RunCommandByScriptContent(string(reqData.ScriptContent))

	data := Reponse{Results: string(scriptOutput)}

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
	var reqData RequestData
	if err := c.ShouldBindJSON(&reqData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("reqData", reqData)
	fmt.Println("reqData.Url", reqData.Url)

	// Get content from url
	responseFromScriptUrl, err := http.Get(reqData.Url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer responseFromScriptUrl.Body.Close()
	fmt.Println("responseFromScriptUrl", responseFromScriptUrl)
	scriptContent, err := ioutil.ReadAll(responseFromScriptUrl.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("string(scriptContent)", string(scriptContent))

	scriptOutput := runCommand.RunCommandByScriptContent(string(scriptContent))
	fmt.Println("scriptOutput", scriptOutput)

	data := Reponse{Results: string(scriptOutput)}

	c.JSON(http.StatusOK, gin.H{"results": data})

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
// @Router /dev/getOriginalMetadataJson [get]
func GetOriginalMetadataJson(c *gin.Context) {
	originalMetadataJson, err := agentOriginMetadataJsonManager.GetOriginalMetadataJson()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"results": originalMetadataJson})
}

func GetAppDataPathByAppName(appName string) string {
	fmt.Println("GetAppDataPathByAppName - start")

	var appDataByAppNamePath string
	switch runtime.GOOS {
	case "linux":
		if os.Getenv("XDG_CONFIG_HOME") != "" {
			appDataByAppNamePath = filepath.Join(os.Getenv("XDG_CONFIG_HOME"), appName)
		} else {
			appDataByAppNamePath = filepath.Join(os.Getenv("HOME"), ".config", appName)
		}
	case "windows":
		appDataByAppNamePath = filepath.Join(os.Getenv("APPDATA"), appName)
	case "darwin":
		appDataByAppNamePath = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", appName)
	default:
		fmt.Println("Unsupported operating system")
		os.Exit(1)
	}

	fmt.Println("GetAppDataPathByAppName:", appDataByAppNamePath)

	return appDataByAppNamePath
}

// @Summary Test
// @Description Test
// @Accept  json
// @Produce  json
// @Router /dev/test [get]
func Test(c *gin.Context) {

	type Config struct {
		Server struct {
			Host string `json:"host"`
			Port int    `json:"port"`
		} `json:"server"`
	}

	appName := "HelloWorldGoAgent"
	fileName := "config.json"

	// Create a new instance of Viper
	v := viper.New()

	// Set the configuration file name
	v.SetConfigFile(fileName)

	// Set the default appData path for Linux, Windows, and macOS systems
	var appDataPath string = GetAppDataPathByAppName(appName)
	configFileLocation := filepath.Join(appDataPath, fileName)

	// Set the configuration file name with the full path
	v.SetConfigFile(configFileLocation)

	// Set some configuration options
	v.Set("server.address", "localhost")
	v.Set("server.port", 8080)

	// Create the configuration directory if it doesn't exist
	if _, err := os.Stat(appDataPath); os.IsNotExist(err) {
		os.MkdirAll(appDataPath, 0755)
	}

	// Create the configuration file if it doesn't exist
	if _, err := os.Stat(configFileLocation); os.IsNotExist(err) {
		// Save the configuration file, create it if it doesn't exist
		err := v.SafeWriteConfig()
		if err != nil {
			fmt.Printf("Error creating config file: %v\n", err)
		}
	} else {
		// Read the configuration file
		err := v.ReadInConfig()
		if err != nil {
			fmt.Printf("Error reading config file: %v\n", err)
		}
	}

	// Save changes to the configuration file
	err := v.WriteConfigAs(configFileLocation)
	if err != nil {
		fmt.Printf("Error writing config file: %v\n", err)
	}

	var config Config
	err = v.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	c.JSON(http.StatusOK, gin.H{"results": config})
}
