package TOS

import (
	"errors"
	"fmt"
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
		originalMetadataPath = "/etc/.helloWorldGoAgent/original_metadata.json"
	} else if runtime.GOOS == "windows" {
		originalMetadataPath = "C:\\Users\\Administrator\\AppData\\Roaming\\.helloWorldGoAgent\\original_metadata.json"
	} else if runtime.GOOS == "darwin" {
		originalMetadataPath = "/Users/jieanyang/Documents/freelancer_work/ansys/HelloWorldGoAgent/src/tools/3_unit/TAgentMetadataManager//original_metadata.json" // @DEV
	} else {
		fmt.Println("Unsupported operating system")
		error := errors.New("unsupported operating system")
		return originalMetadataPath, error
	}

	return originalMetadataPath, nil
}
