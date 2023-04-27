package requestWithBackend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var BACKEND_ENDPOINT string = "https://ff66-2a01-cb16-60-e0c3-a023-d11-c4af-ce6a.ngrok-free.app"

// === GetOperationCommandFromBackend - start ===
func GetOperationCommandFromBackend() *ResponseData {
	responseDataPointer, err := SendGETRequest(BACKEND_ENDPOINT + "/node/aws/getMockOperationCommand")
	if err != nil {
		fmt.Println("Error:", err)
	}

	return responseDataPointer
}

type OperationResult struct {
	ReturnCode int    `json:"returnCode"`
	StdOut     string `json:"stdOut"`
	StdErr     string `json:"stdErr"`
}
type OneOperationCommand struct {
	Id               string          `json:"id"`
	OperationCommand string          `json:"operationCommand"`
	Status           string          `json:"status"`
	OperationScript  string          `json:"operationScript"`
	OperationResult  OperationResult `json:"operationResult"`
	TryTimes         int             `json:"tryTimes"`
}
type ResponseData struct {
	Result OneOperationCommand `json:"result"`
}

func (rd ResponseData) String() string {
	return fmt.Sprintf("===============\nID: %s\nOperation Command: %s\nStatus: %s\nOperation Script: %s\nReturn Code: %d\nStdOut: %s\nStdErr: %s\nTry Times: %d\n===============\n",
		rd.Result.Id,
		rd.Result.OperationCommand,
		rd.Result.Status,
		rd.Result.OperationScript,
		rd.Result.OperationResult.ReturnCode,
		rd.Result.OperationResult.StdOut,
		rd.Result.OperationResult.StdErr,
		rd.Result.TryTimes)
}

func SendGETRequest(url string) (*ResponseData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseData ResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

// === GetOperationCommandFromBackend - end ===

// === PostOperationCommandResultToBackend - start ===

func SendPOSTRequest(url string, data interface{}) (*http.Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func PostOperationCommandResultToBackend(data interface{}) {
	SendPOSTRequest(BACKEND_ENDPOINT+"/node/aws/receiveOperationCommandResult", data)

	// return responseDataPointer
}

// === PostOperationCommandResultToBackend - end ===
