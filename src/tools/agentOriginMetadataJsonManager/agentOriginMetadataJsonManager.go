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
	keyResult, err := GetOriginMetadataJson()
	if err != nil {
		panic(err)
	}
	fmt.Println("keyResult:", keyResult)
}

func GetOriginMetadataJson() (map[string]string, error) {
	os := runtime.GOOS

	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	panic(err)
	// }
	var originMetadataJson map[string]string

	if os == "linux" {
		// originMetadataDir := filepath.Join(homeDir, ".HelloWorldGoAgent")
		originMetadataDir := filepath.Join("/", ".HelloWorldGoAgent")
		originMetadataFile := filepath.Join(originMetadataDir, "origin_metadata.json")

		originMetadataBytes, err := ioutil.ReadFile(originMetadataFile)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(originMetadataBytes, &originMetadataJson)
		if err != nil {
			return nil, err
		}
	} else if os == "windows" {
		// C:\\Users\\Administrator\\.HelloWorldGoAgent\\origin_metadata.json
		// originMetadataDir := filepath.Join(homeDir, ".HelloWorldGoAgent")
		originMetadataDir := filepath.Join("C:\\Users\\Administrator", ".HelloWorldGoAgent")
		originMetadataFile := filepath.Join(originMetadataDir, "origin_metadata.json")

		originMetadataBytes, err := ioutil.ReadFile(originMetadataFile)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(originMetadataBytes, &originMetadataJson)
		if err != nil {
			return nil, err
		}
	}

	return originMetadataJson, nil
}
