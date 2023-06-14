//
// Copyright (C) 2023 ANSYS, Inc. Unauthorized use, distribution, or duplication is prohibited.
//

package agent

import (
	"errors"
	"fmt"
	"os"
	"time"

	"AnsysCSPAgent/src/tools/TAgentMetadataManager"
	"AnsysCSPAgent/src/tools/TRunCommand"
	"AnsysCSPAgent/src/tools/common/TOperationCommand"
	"AnsysCSPAgent/src/tools/requestWithBackend"
)

const (
	Waiting = iota
	Running
)

var ErrWrongState = errors.New("can't take the operation in the current state")

type Agent struct {
	state int
}

func NewAgent() *Agent {
	agent := Agent{
		state: Waiting,
	}

	return &agent
}

// === Launch Agent - start ===
func (agent *Agent) Launch() error {
	fmt.Println("Agent start func - start")

	if agent.state != Waiting {
		return ErrWrongState
	}

	agent.state = Running
	fmt.Println("Start - agent", agent.state)

	// Init agent config.json file
	agent.Init()

	// Get new OperationCommand, execute then send back result
	agent.RunPeriodicTask() // @Prod, block here

	fmt.Println("Agent start func - end")

	return nil
}

// === Launch Agent - end ===

func (agent *Agent) Init() {
	// === load metadata config - start ===

	originalMetadata, err := TAgentMetadataManager.GetOriginalMetadataFileContent()
	if err != nil {
		fmt.Println("Error reading original metadata file:", err.Error())
		os.Exit(1) // @PROD
	}

	// save metadata as an app config file
	TAgentMetadataManager.GetOrCreateConfigFile(originalMetadata)

	// === load metadata config - end ===
}

// === Business logic - polling backend - start ===
func RunPeriodicTask_handleBackendRequestErrorCase(paraErrorCount int, ticker *time.Ticker) {
	switch {
	case paraErrorCount >= 24:
		fmt.Println("RunPeriodicTask_handleBackendRequestErrorCase - Error: set to 15 mins")
		ticker.Reset(15 * time.Minute)
	case paraErrorCount >= 20:
		fmt.Println("RunPeriodicTask_handleBackendRequestErrorCase - Error: set to 5 mins")
		ticker.Reset(5 * time.Minute)
	case paraErrorCount >= 16:
		fmt.Println("RunPeriodicTask_handleBackendRequestErrorCase - Error: set to 1 mins")
		ticker.Reset(1 * time.Minute)
	case paraErrorCount >= 12:
		fmt.Println("RunPeriodicTask_handleBackendRequestErrorCase - Error: set to 30 seconds")
		ticker.Reset(30 * time.Second)
	case paraErrorCount >= 8:
		fmt.Println("RunPeriodicTask_handleBackendRequestErrorCase - Error: set to 15 seconds")
		ticker.Reset(15 * time.Second)
	case paraErrorCount >= 4:
		fmt.Println("RunPeriodicTask_handleBackendRequestErrorCase - Error: set to 10 seconds")
		ticker.Reset(10 * time.Second)
	}
}
func (agent *Agent) RunPeriodicTask() {
	fmt.Println("SetUp for periodic task - start")

	interval := 5 * time.Second
	ticker := time.NewTicker(interval)

	errorCount := 0

	for range ticker.C {
		fmt.Println("RunPeriodicTask - start")

		fmt.Println("for errorCount:", errorCount)

		// Get new OperationCommand script from backend
		responseData, err := requestWithBackend.GetOperationCommandFromBackend()
		if err != nil {
			fmt.Println("RunPeriodicTask - Error:", err)

			errorCount++
			RunPeriodicTask_handleBackendRequestErrorCase(errorCount, ticker)

			continue
		} else {
			fmt.Println("RunPeriodicTask - Successful GetOperationCommandFromBackend - reset ticker to 5 seconds")
			errorCount = 0
			ticker.Reset(5 * time.Second)
		}
		operationScript := responseData.Result.OperationScript
		fmt.Println("RunPeriodicTask - operationScript:", operationScript)

		// execute then send back result
		stdOut, err := TRunCommand.RunCommandByScriptContent(operationScript)
		var returnError bool = false
		var stdErr string = ""
		if err != nil {
			fmt.Println("RunPeriodicTask - RunCommandByScriptContent - Error:", err)
			stdErr = stdOut + "\n======\n" + err.Error()
			stdOut = ""
			returnError = true
		}

		postResult := TOperationCommand.OneOperationCommand{
			Id:               responseData.Result.Id,
			OperationCommand: responseData.Result.OperationCommand,
			Status:           responseData.Result.Status,
			OperationScript:  responseData.Result.OperationScript,
			OperationResult: TOperationCommand.OperationResult{
				StdOut:      stdOut,
				StdErr:      stdErr,
				ReturnError: returnError,
			},
			TryTimes: responseData.Result.TryTimes,
		}

		requestWithBackend.PostOperationCommandResultToBackend(postResult)
		fmt.Println("RunPeriodicTask - end")
	}

	fmt.Println("SetUp for periodic task - end")
}

// === Business logic - polling backend - end ===
