//
// Copyright (C) 2023 ANSYS, Inc. Unauthorized use, distribution, or duplication is prohibited.
//

package requestWithBackend

import (
	"AnsysCSPAgent/src/tools/TOperationCommand"
	"AnsysCSPAgent/src/tools/common/TRequest"
	"fmt"
)

var BACKEND_ENDPOINT string = "https://0822-2a01-cb08-ad0-f700-b885-c4bc-fd0c-bc93.ngrok-free.app"

// === GetOperationCommandFromBackend - start ===
func GetOperationCommandFromBackend() (*TOperationCommand.OneOperationCommandResponseData, error) {
	responseDataPointer, err := TRequest.SendGETRequest(BACKEND_ENDPOINT + "/node/aws/getMockOperationCommand")
	if err != nil {
		fmt.Println("GetOperationCommandFromBackend - Error:", err)
	}

	return responseDataPointer, err
}

// === GetOperationCommandFromBackend - end ===

// === PostOperationCommandResultToBackend - start ===

func PostOperationCommandResultToBackend(data interface{}) {
	TRequest.SendPOSTRequest(BACKEND_ENDPOINT+"/node/aws/receiveOperationCommandResult", data)

	// return responseDataPointer
}

// === PostOperationCommandResultToBackend - end ===
