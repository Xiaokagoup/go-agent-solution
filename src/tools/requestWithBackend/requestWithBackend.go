package requestWithBackend

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var BACKEND_ENDPOINT string = "https://e011-90-3-247-18.ngrok-free.app"

// === GetOperationCommandFromBackend - start ===
func GetOperationCommandFromBackend() (*ResponseData, error) {
	responseDataPointer, err := SendGETRequest(BACKEND_ENDPOINT + "/node/aws/getMockOperationCommand")
	if err != nil {
		fmt.Println("Error:", err)
	}

	return responseDataPointer, err
}

type OperationResult struct {
	StdOut      string `json:"stdOut"`
	StdErr      string `json:"stdErr"`
	ReturnError bool   `json:"returnError"`
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
	return fmt.Sprintf("===============\nID: %s\nOperation Command: %s\nStatus: %s\nOperation Script: %s\nStdOut: %s\nStdErr: %s\nReturn Error: %t\nTry Times: %d\n===============\n",
		rd.Result.Id,
		rd.Result.OperationCommand,
		rd.Result.Status,
		rd.Result.OperationScript,
		rd.Result.OperationResult.StdOut,
		rd.Result.OperationResult.StdErr,
		rd.Result.OperationResult.ReturnError,
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
		fmt.Println("ioutil.ReadAll Error:", err)
		return nil, err
	}

	var responseData ResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("json.Unmarshal Error:", err)
		var newErr error = errors.New("The data returned from the backend is incorrect.\n" + err.Error())
		return nil, newErr
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
