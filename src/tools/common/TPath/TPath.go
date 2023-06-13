//
// Copyright (C) 2023 ANSYS, Inc. Unauthorized use, distribution, or duplication is prohibited.
//

package TPath

import (
	"errors"
	"fmt"
	"os"
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
	} else {
		fmt.Println("Unsupported operating system")
		error := errors.New("unsupported operating system")
		return originalMetadataPath, error
	}

	return originalMetadataPath, nil
}
