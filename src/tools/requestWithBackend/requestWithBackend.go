package requestWithBackend

import (
	"AnsysCSPAgent/src/tools/4_base/TRequest"
	"fmt"
)

var BACKEND_ENDPOINT string = "https://e011-90-3-247-18.ngrok-free.app"

// === GetOperationCommandFromBackend - start ===
func GetOperationCommandFromBackend() (*TRequest.ResponseData, error) {
	responseDataPointer, err := TRequest.SendGETRequest(BACKEND_ENDPOINT + "/node/aws/getMockOperationCommand")
	if err != nil {
		fmt.Println("Error:", err)
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
