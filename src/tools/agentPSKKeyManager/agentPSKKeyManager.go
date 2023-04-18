package agentPSKKeyManager

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("run main in agentKeysManager.go")
	keyResult := GetPSKKey()
	fmt.Println("keyResult:", keyResult)
}

func GetPSKKey() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	keyDir := filepath.Join(homeDir, ".HelloWorldGoAgent")
	keyFile := filepath.Join(keyDir, "PSK_key.txt")

	keyBytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		panic(err)
	}

	key := string(keyBytes)
	fmt.Println("Key:", key)

	return key
}
