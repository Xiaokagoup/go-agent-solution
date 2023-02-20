package agentKeysManager

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const keyFileName = "keys.json"

const (
	AuthKey = iota
	TransferKey
	SessionKey
)

func getKeyType(key Key) string {
	switch key.Type {
	case AuthKey:
		return "AuthKey"
	case TransferKey:
		return "TransferKey"
	case SessionKey:
		return "SessionKey"
	default:
		return "Unknown Key Type"
	}
}

type Key struct {
	Type      int
	Value     string
	ExpiresAt time.Time
}

func main() {
	fmt.Println("run main in agentKeysManager.go")

	currentKeys := []Key{}

	firstKey := generateKey(AuthKey)
	fmt.Println("firstKey", firstKey)
	fmt.Printf("firstKey type %T\n", firstKey)

	currentKeys = append(currentKeys, firstKey)
	fmt.Println("currentKeys", currentKeys)

	saveKeysInFile(currentKeys, keyFileName)

	currentKeysFromFile, _ := loadKeys(keyFileName)
	fmt.Println("currentKeysFromFile", currentKeysFromFile)

	secondKey := generateKey(AuthKey)
	fmt.Println("secondKey", secondKey)
	fmt.Printf("secondKey type %T\n", secondKey)

	nextKeys := addKey(currentKeysFromFile, secondKey)
	fmt.Println("nextKeys", nextKeys)

	fmt.Println("====================================")

	currentKeys, _ = saveKeysInFile(nextKeys, keyFileName)
	fmt.Println("before deleteKey", len(currentKeys))
	currentKeysAfterDelete := deleteKey(currentKeys, firstKey)
	fmt.Println("currentKeysAfterDelete", len(currentKeysAfterDelete))
	currentKeys, _ = saveKeysInFile(currentKeysAfterDelete, keyFileName)

	fmt.Println("====================================")

}

func loadKeys(keyFileName string) ([]Key, error) {
	// Read the keys from the key file
	data, err := ioutil.ReadFile(keyFileName)
	if err != nil {
		return nil, fmt.Errorf("Error reading keys from file: %v", err)
	}

	// Parsethe keys from the data
	var keys []Key
	err = json.Unmarshal(data, &keys)

	return keys, nil
}

func containKey(currentKeys []Key, key Key) bool {
	for _, k := range currentKeys {
		if k.Value == key.Value {
			return true
		}
	}
	return false
}

func generateKey(keyType int) Key {
	// Generate a random key of 32 bytes
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("Error generating key:", err)
		os.Exit(1)
	}
	return Key{
		Type:      keyType,
		Value:     base64.StdEncoding.EncodeToString(key),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}
}

func addKey(currentKeys []Key, addedKey Key) []Key {
	if !containKey(currentKeys, addedKey) {
		nextKeys := append(currentKeys, addedKey)
		return nextKeys
	}
	return currentKeys
}

func deleteKey(currentKeys []Key, deletedKey Key) []Key {
	isEqualKey := func(a Key, b Key) bool {
		return a.Type == b.Type && a.Value == b.Value
	}
	for i, k := range currentKeys {
		fmt.Println("i", i)
		fmt.Println("k", k)
		fmt.Println("isEqualKey(k, deletedKey)", isEqualKey(k, deletedKey))
		if isEqualKey(k, deletedKey) {
			nextKeys := append(currentKeys[:i], currentKeys[i+1:]...)
			return nextKeys
		}
	}
	return currentKeys
}

func saveKeysInFile(newKeys []Key, keyFileName string) ([]Key, error) {
	// Convert the keys to JSON
	data, err := json.Marshal(newKeys)
	if err != nil {
		return nil, fmt.Errorf("Error converting keys to JSON: %v", err)
	}

	// Write the keys to the key file
	err = ioutil.WriteFile(keyFileName, data, 0644)
	if err != nil {
		return nil, fmt.Errorf("Error writing keys to the file: %v", err)
	}

	return newKeys, nil
}
