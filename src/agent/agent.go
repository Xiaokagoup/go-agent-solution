package agent

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	agtHttp "AnsysCSPAgent/src/agentHttp"
	"AnsysCSPAgent/src/tools/3_unit/TAgentMetadataManager"
	"AnsysCSPAgent/src/tools/3_unit/TOperationCommand"
	"AnsysCSPAgent/src/tools/4_base/TOS"
	"AnsysCSPAgent/src/tools/TRunCommand"
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

type Original_Metadata struct {
	PSK_Key string `json:"psk_key"`
}

func (agent *Agent) Init() {
	// === load metadata config - start ===

	// Get original metadata path
	originalMetadataPath, err := TOS.GetAgentOriginalMetadataFilePath()
	if err != nil {
		fmt.Println("Error getting original metadata path:", err.Error())
		os.Exit(1) // @PROD
	}

	// Read the JSON file
	originalMetaData, err := ioutil.ReadFile(originalMetadataPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Unmarshal the JSON data into a slice of Person structs
	var metadata Original_Metadata
	err = json.Unmarshal(originalMetaData, &metadata)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println("metadata found", metadata, metadata.PSK_Key)

	// save metadata as an app config file
	TAgentMetadataManager.GetOrCreateConfigFileWithSpecifiedPskKey(metadata.PSK_Key)

	// === load metadata config - end ===
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

	// Start web interface for development
	go agtHttp.StartHttp() // @DEV

	// Get new OperationCommand, execute then send back result
	agent.RunPeriodicTask() // @Prod, block here

	fmt.Println("Agent start func - end")

	return nil
}

// === Launch Agent - end ===

// === Business logic - polling backend - start ===
func (agent *Agent) RunPeriodicTask() {
	fmt.Println("SetUp for periodic task - start")

	defaultInterval := 5 * time.Second
	ticker := time.NewTicker(defaultInterval)
	intervalChan := make(chan time.Duration)

	for {
		select {
		case <-ticker.C:
			fmt.Println("RunPeriodicTask - start")

			// Get new OperationCommand script from backend
			responseData, err := requestWithBackend.GetOperationCommandFromBackend()
			if err != nil {
				fmt.Println("RunPeriodicTask - Error:", err)
				continue
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

		case newInterval := <-intervalChan:
			fmt.Println("RunPeriodicTask - reset new newInterval:", newInterval)
			ticker.Stop()
			ticker = time.NewTicker(newInterval)
		}
	}

	fmt.Println("SetUp for periodic task - end")
}

// === Business logic - polling backend - end ===
