package TRequest

import (
	"AnsysCSPAgent/src/tools/3_unit/TOperationCommand"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type StringRequestData struct {
	Result string `json:"result"`
}

func SendGETRequest(url string) (*TOperationCommand.OneOperationCommandResponseData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("SendGETRequest - ioutil.ReadAll - Error:", err)
		return nil, err
	}

	var responseData TOperationCommand.OneOperationCommandResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("SendGETRequest - json.Unmarshal - Error:", err)
		var newErr error = errors.New("The data returned from the backend is incorrect. " + err.Error())
		return nil, newErr
	}

	return &responseData, nil
}

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
