//
// Copyright (C) 2023 ANSYS, Inc. Unauthorized use, distribution, or duplication is prohibited.
//

package requestWithBackend

import (
	"AnsysCSPAgent/src/tools/common/TOperationCommand"
	"AnsysCSPAgent/src/tools/common/TRequest"
	"fmt"
	"time"
)

// var BACKEND_ENDPOINT string = "http://ec2-3-121-159-217.eu-central-1.compute.amazonaws.com/request/jieayang"
var BACKEND_ENDPOINT string = "http://localhost:8080/node"

// === GetOperationCommandFromBackend - start ===
func GetOperationCommandFromBackend() (*TOperationCommand.OneOperationCommandResponseData, error) {
	responseChan := make(chan *TOperationCommand.OneOperationCommandResponseData)
	errorChan := make(chan error)

	// Start a goroutine to perform the HTTP request
	go func() {
		responseDataPointer, err := TRequest.SendGETRequest(BACKEND_ENDPOINT + "/v2/agent/getOperationCommand")
		if err != nil {
			errorChan <- err
			return
		}
		responseChan <- responseDataPointer
	}()

	// Wait for either the response or an error, or timeout after 10 seconds
	select {
	case responseData := <-responseChan:
		return responseData, nil
	case err := <-errorChan:
		fmt.Println("GetOperationCommandFromBackend - Error:", err)
		return nil, err
	case <-time.After(10 * time.Second): // The request is taking too long, cancel it
		return nil, fmt.Errorf("timeout while waiting for backend response")
	}

}

// === GetOperationCommandFromBackend - end ===

// === PostOperationCommandResultToBackend - start ===

func PostOperationCommandResultToBackend(data interface{}) {
	TRequest.SendPOSTRequest(BACKEND_ENDPOINT+"/v2/agent/receiveOperationCommandResult", data)

	// return responseDataPointer
}

// === PostOperationCommandResultToBackend - end ===
