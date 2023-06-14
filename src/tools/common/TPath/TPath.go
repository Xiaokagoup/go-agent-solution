//
// Copyright (C) 2023 ANSYS, Inc. Unauthorized use, distribution, or duplication is prohibited.
//

package TPath

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetAgentOriginalMetadataFilePath() (string, error) {
	var originalMetadataPath string = ""

	if runtime.GOOS == "linux" {
		originalMetadataPath = "/etc/.ansysCSPAgent/original_metadata.json"
	} else if runtime.GOOS == "windows" {
		originalMetadataPath = "C:\\Users\\Administrator\\AppData\\Roaming\\.ansysCSPAgent\\original_metadata.json"

		// Check if file exists
		_, err := os.Stat(originalMetadataPath)
		if os.IsNotExist(err) {
			// File does not exist, set an alternative path
			originalMetadataPath = "C:\\Users\\Administrator_ansys\\AppData\\Roaming\\.ansysCSPAgent\\original_metadata.json"
		}
	} else if runtime.GOOS == "darwin" {
		originalMetadataPath = "/Users/jieanyang/Documents/freelancer_work/ansys/HelloWorldGoAgent/src/tools/TAgentMetadataManager/metadata.json" // @DEV
	} else {
		fmt.Println("Unsupported operating system")
		error := errors.New("unsupported operating system")
		return originalMetadataPath, error
	}

	return originalMetadataPath, nil
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
