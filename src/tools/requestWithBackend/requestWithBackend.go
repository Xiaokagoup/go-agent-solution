//
// Copyright (C) 2023 ANSYS, Inc. Unauthorized use, distribution, or duplication is prohibited.
//

package requestWithBackend

import (
	"AnsysCSPAgent/src/tools/common/TOperationCommand"
	"AnsysCSPAgent/src/tools/common/TRequest"
	"fmt"
)

var BACKEND_ENDPOINT string = "http://ec2-3-121-159-217.eu-central-1.compute.amazonaws.com/request/jieayang"

// === GetOperationCommandFromBackend - start ===
func GetOperationCommandFromBackend() (*TOperationCommand.OneOperationCommandResponseData, error) {
	responseDataPointer, err := TRequest.SendGETRequest(BACKEND_ENDPOINT + "/v2/agent/getOperationCommand")
	if err != nil {
		fmt.Println("GetOperationCommandFromBackend - Error:", err)
	}

	return responseDataPointer, err
}

// === GetOperationCommandFromBackend - end ===

// === PostOperationCommandResultToBackend - start ===

func PostOperationCommandResultToBackend(data interface{}) {
	TRequest.SendPOSTRequest(BACKEND_ENDPOINT+"/v2/agent/receiveOperationCommandResult", data)

	// return responseDataPointer
}

// === PostOperationCommandResultToBackend - end ===
