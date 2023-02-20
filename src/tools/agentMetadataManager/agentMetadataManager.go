package agentJsonManager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Metadata struct {
	ClientId      string `json:"clientId"`
	CloudProvider string `json:"cloudProvider"`
	Region        string `json:"region"`
	NodeType      string `json:"nodeType"`
	CreatedAt     string `json:"createdAt"`
}

func main() {
	fileName := "metadata.json"

	metadatas := Metadata{
		ClientId:      "3",
		CloudProvider: "AWS",
		Region:        "eu-west-3",
		NodeType:      "t2.micro",
		CreatedAt:     "2023-02-16T11:22:35.040Z",
	}

	writeMetadataToFile(fileName, metadatas)

	fileContent, _ := readMetadataFromFile(fileName)
	fmt.Printf("fileContent %T, %v \n", fileContent, fileContent)
	fmt.Println("fileContent", *fileContent)

}

func readMetadataFromFile(fileName string) (*Metadata, error) {
	// Read the contents of the file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data in to a slice of User structs
	var metadata *Metadata
	err = json.Unmarshal(data, &metadata)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}

func writeMetadataToFile(fileName string, metadata Metadata) error {
	// Convert the slice of Usser structs to JSON data
	data, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
