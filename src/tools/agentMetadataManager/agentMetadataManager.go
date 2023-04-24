package agentMetadataManager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
}

func main() {
	// Test function
	fileName := "metadata.json"

	metadatas := Metadata{
		ClientId:      "3",
		CloudProvider: "AWS",
		Region:        "eu-west-3",
		NodeType:      "t2.micro",
		CreatedAt:     "2023-02-16T11:22:35.040Z",
	}

	writeMetadataToFile(fileName, metadatas)

	fileContent, _ := readMetadataFromFile(fileName)
	fmt.Printf("fileContent %T, %v \n", fileContent, fileContent)
	fmt.Println("fileContent", *fileContent)

}

func readMetadataFromFile(fileName string) (*Metadata, error) {
	// Read the contents of the file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data in to a slice of User structs
	var metadata *Metadata
	err = json.Unmarshal(data, &metadata)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}

func writeMetadataToFile(fileName string, metadata Metadata) error {
	// Convert the slice of Usser structs to JSON data
	data, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
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

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"server"`
	ConfigFileLocation string `json:"configFileLocation"`
}

// To Custimize the output of the struct Config when printing it
func (c Config) String() string {
	return fmt.Sprintf("Host: %s, Port: %d", c.Server.Host, c.Server.Port)
}

func GetOrCreateConfigFile() Config {
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
	v.Set("configFileLocation", configFileLocation)

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

	return config
}
