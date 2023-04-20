package agentOriginMetadataJsonManager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println("run main in agentKeysManager.go")
	keyResult, err := GetOriginalMetadataJson()
	if err != nil {
		panic(err)
	}
	fmt.Println("keyResult:", keyResult)
}

func GetOriginalMetadataJson() (map[string]string, error) {
	os := runtime.GOOS

	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	panic(err)
	// }
	var originalMetadataJson map[string]string

	if os == "linux" {
		// originalMetadataDir := filepath.Join(homeDir, ".HelloWorldGoAgent")
		originalMetadataDir := filepath.Join("/", ".HelloWorldGoAgent")
		originalMetadataFile := filepath.Join(originalMetadataDir, "original_metadata.json")

		originalMetadataBytes, err := ioutil.ReadFile(originalMetadataFile)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(originalMetadataBytes, &originalMetadataJson)
		if err != nil {
			return nil, err
		}
	} else if os == "windows" {
		// C:\\Users\\Administrator\\.HelloWorldGoAgent\\original_metadata.json
		// originalMetadataDir := filepath.Join(homeDir, ".HelloWorldGoAgent")
		originalMetadataDir := filepath.Join("C:\\Users\\Administrator", ".HelloWorldGoAgent")
		originalMetadataFile := filepath.Join(originalMetadataDir, "original_metadata.json")

		originalMetadataBytes, err := ioutil.ReadFile(originalMetadataFile)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(originalMetadataBytes, &originalMetadataJson)
		if err != nil {
			return nil, err
		}
	}

	return originalMetadataJson, nil
}
