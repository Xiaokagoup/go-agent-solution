//
// Copyright (C) 2023 ANSYS, Inc. Unauthorized use, distribution, or duplication is prohibited.
//

package TAgentMetadataManager

import (
	"AnsysCSPAgent/src/tools/common/TPath"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/viper"
)

func main() {
	// Test function
	fileName := "metadata.json"

	metadatas := Metadata{
		ClientId:      "4",
		CloudProvider: "AWS",
		Region:        "eu-west-3",
		NodeType:      "t2.micro",
		CreatedAt:     "2023-02-16T11:22:35.040Z",
		PSK_Key:       "12345678901234567890123456789012",
	}

	writeMetadataToFile(fileName, &metadatas)

	fileContent, _ := ReadMetadataFromFile(fileName)

	fmt.Printf("Metadata: %+v\n", fileContent)
}

func ReadMetadataFromFile(filePath string) (*Metadata, error) {
	// Read the contents of the file
	data, err := ioutil.ReadFile(filePath)
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

func writeMetadataToFile(filePath string, metadata *Metadata) (*Metadata, error) {
	// Convert the slice of Usser structs to JSON data
	data, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}

	// Write the JSON data to the file
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}

type Metadata struct {
	ClientId      string `json:"clientId"`
	CloudProvider string `json:"cloudProvider"`
	Region        string `json:"region"`
	NodeType      string `json:"nodeType"`
	CreatedAt     string `json:"createdAt"`
	PSK_Key       string `json:"psk_key"`
}

func (m Metadata) String() string {
	return fmt.Sprintf("ClientId: %s, CloudProvider: %s, Region: %s, NodeType: %s, CreatedAt: %s, PSK_Key: %s",
		m.ClientId, m.CloudProvider, m.Region, m.NodeType, m.CreatedAt, m.PSK_Key)
}

// type ConfigResponseData struct {
// 	Result Config `json:"result"`
// }

// type Config struct {
// 	Server struct {
// 		Host string `json:"host"`
// 		Port int    `json:"port"`
// 	} `json:"server"`
// 	ConfigFileLocation string `json:"configFileLocation"`
// 	PSK_Key            string `json:"psk_key"`
// }

// To Custimize the output of the struct Config when printing it
// func (c Config) String() string {
// 	return fmt.Sprintf("Host: %s, Port: %d", c.Server.Host, c.Server.Port)
// }

func GetOriginalMetadataFileContent() (*Metadata, error) {
	osServiceManagerAppName := "ansysCSPAgentManagerService"
	agentAppName := "ansysCSPAgent"
	fileName := "config.json"

	// Create a new instance of Viper
	v := viper.New()

	// Set the configuration file name
	v.SetConfigFile(fileName)

	// Set the default appData path for Linux, Windows, and macOS systems
	var agentAppDataPath string = TPath.GetAgentAppDataPathByAppName(osServiceManagerAppName, agentAppName)
	configFileLocation := filepath.Join(agentAppDataPath, fileName)

	// Read the configuration file
	fileContent, err := ReadMetadataFromFile(configFileLocation)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return nil, err // empty Metadata object
	}

	return fileContent, nil
}

func GetOrCreateConfigFile(metaData *Metadata) (*Metadata, error) {
	osServiceManagerAppName := "ansysCSPAgentManagerService"
	agentAppName := "ansysCSPAgent"
	fileName := "config.json"

	// Set the default appData path for Linux, Windows, and macOS systems
	var agentAppDataPath string = TPath.GetAgentAppDataPathByAppName(osServiceManagerAppName, agentAppName)
	configFileLocation := filepath.Join(agentAppDataPath, fileName)

	// Create or rewrite config.json file
	if metaData != nil {
		metaData, err := writeMetadataToFile(configFileLocation, metaData)
		if err != nil {
			fmt.Printf("Error creating or rewriting config file: %v\n", err)
			return nil, err // empty Metadata object
		}
		return metaData, nil
	}

	// Read the configuration file
	fileContent, err := ReadMetadataFromFile(configFileLocation)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return nil, err // empty Metadata object
	}

	return fileContent, nil
}
