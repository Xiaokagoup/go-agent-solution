package agentKeysManager

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
)

const keyFile = "keys.bin"

func generateKey() []byte {
	// Generate a random key of 32 bytes
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("Error generating key:", err)
		os.Exit(1)
	}
	return key
}

func generateKeys(num int) [][]byte {
	// Generate [num] random keys of 32 bytes each
	keys := make([][]byte, num)
	for i := range keys {
		key := generateKey()
		keys[i] = key
	}
	return keys
}

func addKeyArray(keys [][]byte, key []byte) [][]byte {
	return append(keys, key)
}

func deleteKeyArray(keys [][]byte, index int) [][]byte {
	return append(keys[:index], keys[index+1:]...)
}

func loadKeys(keyFile string) ([][]byte, error) {
	// Read the keys from the key file
	data, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("Error reading keys from file: %v", err)
	}

	// Parsethe keys from the data
	keys := [][]byte{}
	for i := 0; i < len(data); i += 32 {
		key := data[i : i+32]
		keys = append(keys, key)
	}
	return keys, nil
}

func containKey(keys [][]byte, key []byte, keyFile string) bool {
	for _, k := range keys {
		if bytes.Equal(k, key) {
			return true
		}
	}
	return false
}

func saveKeys(newKeys [][]byte, keyFile string) error {
	// // Load the existing keys from the file
	// existingKeys, err := loadKeys(keyFile)
	// if err != nil {
	// 	return err
	// }

	// // Add the new keys to the existing keys
	// for _, oneNewKey := range newKeys {
	// 	// Check if the key already exists
	// 	if !containsKey(existingKeys, oneNewKey,key) {
	// 		existingKeys = agent
	// 	}
	// }

	// write the keys to the key file
	data := []byte{}
	for _, key := range newKeys {
		data = append(data, key...)
	}
	err := ioutil.WriteFile(keyFile, data, 0644)
	if err != nil {
		return fmt.Errorf("Error writing keys to the file: %v", err)
	}
	return nil
}

func saveNewKeysInFile(key []byte, keyFile string) error {
	// load the existing keys from the file
	keys, err := loadKeys(keyFile)
	if err != nil {
		return err
	}

	// Check if the key already exists
	for _, k := range keys {
		if bytes.Equal(k, key) {
			return nil
		}
	}

	// Add the new key
	keys = append(keys, key)

	// Save the modified keys back to the file
	return saveKeys(keys, keyFile)
}
