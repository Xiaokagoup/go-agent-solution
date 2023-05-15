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

type ResponseData struct {
	Result TOperationCommand.OneOperationCommand `json:"result"`
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
