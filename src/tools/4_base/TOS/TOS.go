package TOS

import (
	"errors"
	"fmt"
	"os"
	"runtime"
)

const (
	Linux = iota
	Windows
	MacOS
	Unknown
)

func GetOSName() int {
	os := runtime.GOOS
	if os == "windows" {
		fmt.Println("Windows operating system detected")
		return Windows
	} else if os == "linux" {
		fmt.Println("Linux operating system detected")
		return Linux
	} else if os == "darwin" {
		fmt.Println("Mac operating system detected")
		return MacOS
	}
	fmt.Printf("Unknown operating system: %s\n", os)
	return Unknown
}

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
	} else {
		fmt.Println("Unsupported operating system")
		error := errors.New("unsupported operating system")
		return originalMetadataPath, error
	}

	return originalMetadataPath, nil
}
