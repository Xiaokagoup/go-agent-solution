//
// Copyright (C) 2023 ANSYS, Inc. Unauthorized use, distribution, or duplication is prohibited.
//

package TAgentMetadataManager

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Metadata struct {
	ClientId      string `json:"clientId"`
	CloudProvider string `json:"cloudProvider"`
	Region        string `json:"region"`
	NodeType      string `json:"nodeType"`
	CreatedAt     string `json:"createdAt"`
	PSK_Key       string `json:"psk_key"`
}

func GetAgentAppDataPathByAppName(osServiceManagerAppName string, agentAppName string) string {
	fmt.Println("GetAppDataPathByAppName - start")

	var appDataByAppNamePath string
	switch runtime.GOOS {
	case "linux":
		appDataByAppNamePath = filepath.Join("/usr/local/go", osServiceManagerAppName, agentAppName)
	case "windows":
		appDataByAppNamePath = filepath.Join("C:\\go", osServiceManagerAppName, agentAppName)
	case "darwin":
		appDataByAppNamePath = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", osServiceManagerAppName, agentAppName)
	default:
		fmt.Println("Unsupported operating system")
		os.Exit(1)
	}

	fmt.Println("GetAppDataPathByAppName:", appDataByAppNamePath)

	return appDataByAppNamePath
}

type ConfigResponseData struct {
	Result Config `json:"result"`
}

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"server"`
	ConfigFileLocation string `json:"configFileLocation"`
	PSK_Key            string `json:"psk_key"`
}

// To Custimize the output of the struct Config when printing it
func (c Config) String() string {
	return fmt.Sprintf("Host: %s, Port: %d", c.Server.Host, c.Server.Port)
}

func GetConfigFileContent() (Config, error) {
	osServiceManagerAppName := "ansysCSPAgentManagerService"
	agentAppName := "ansysCSPAgent"
	fileName := "config.json"

	// Create a new instance of Viper
	v := viper.New()

	// Set the configuration file name
	v.SetConfigFile(fileName)

	// Set the default appData path for Linux, Windows, and macOS systems
	var agentAppDataPath string = GetAgentAppDataPathByAppName(osServiceManagerAppName, agentAppName)
	configFileLocation := filepath.Join(agentAppDataPath, fileName)

	// Set the configuration file name with the full path
	v.SetConfigFile(configFileLocation)
	fmt.Println("configFileLocation:", configFileLocation)

	// Read the configuration file
	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return Config{}, err
	}

	var config Config
	err = v.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	return config, nil
}

func GetOrCreateConfigFileWithSpecifiedPskKey(pskKey string) Config {
	osServiceManagerAppName := "ansysCSPAgentManagerService"
	agentAppName := "ansysCSPAgent"
	fileName := "config.json"

	// Create a new instance of Viper
	v := viper.New()

	// Set the configuration file name
	v.SetConfigFile(fileName)

	// Set the default appData path for Linux, Windows, and macOS systems
	var agentAppDataPath string = GetAgentAppDataPathByAppName(osServiceManagerAppName, agentAppName)
	configFileLocation := filepath.Join(agentAppDataPath, fileName)

	// Set the configuration file name with the full path
	v.SetConfigFile(configFileLocation)

	// Set some configuration options
	v.Set("server.address", "localhost")
	v.Set("server.port", 8080)
	v.Set("configFileLocation", configFileLocation)
	v.Set("psk_key", pskKey)

	// Create the configuration directory if it doesn't exist
	if _, err := os.Stat(agentAppDataPath); os.IsNotExist(err) {
		os.MkdirAll(agentAppDataPath, 0755)
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

	return config
}
