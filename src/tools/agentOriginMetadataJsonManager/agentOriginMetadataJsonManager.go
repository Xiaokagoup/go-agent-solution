package agentOriginMetadataJsonManager

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	fmt.Println("run main in agentKeysManager.go")
	keyResult := GetOriginMetadataJson()
	fmt.Println("keyResult:", keyResult)
}

func GetOriginMetadataJson() string {
	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	panic(err)
	// }

	// keyDir := filepath.Join(homeDir, ".HelloWorldGoAgent")
	keyDir := filepath.Join("/", ".HelloWorldGoAgent")
	keyFile := filepath.Join(keyDir, "origin_metadata.json")

	keyBytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		panic(err)
	}

	key := string(keyBytes)
	fmt.Println("Key:", key)

	return key
}
