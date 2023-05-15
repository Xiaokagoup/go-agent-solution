package agent

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"time"

	agtHttp "AnsysCSPAgent/src/agentHttp"
	"AnsysCSPAgent/src/tools/4_base/TRequest"
	"AnsysCSPAgent/src/tools/agentMetadataManager"
	"AnsysCSPAgent/src/tools/requestWithBackend"
	"AnsysCSPAgent/src/tools/runCommand"
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
	// load config
	// load modules
	// load heartbeat signal
	// load metrics
	// load message

	var originalMetadataPath string
	if runtime.GOOS == "linux" {
		originalMetadataPath = "/etc/.helloWorldGoAgent/original_metadata.json"
	} else if runtime.GOOS == "windows" {
		originalMetadataPath = "C:\\Users\\Administrator\\AppData\\Roaming\\.helloWorldGoAgent\\original_metadata.json"
	} else {
		fmt.Println("Unsupported operating system")
		return     // @DEV
		os.Exit(1) // @PROD
	}

	// originalMetadataPath = "/Users/jieanyang/Documents/freelancer_work/ansys/HelloWorldGoAgent/src/tools/agentMetadataManager/original_metadata.json" // @DEV

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

	agentMetadataManager.GetOrCreateConfigFileWithSpecifiedPskKey(metadata.PSK_Key) // save psk key to config file

}

// Business logic
func (agent *Agent) Start() error {
	fmt.Println("Agent start func - start")

	if agent.state != Waiting {
		return ErrWrongState
	}

	agent.state = Running
	fmt.Println("Start - agent", agent.state)

	agent.Init()
	go agtHttp.StartHttp()  // @Prod
	agent.RunPeriodicTask() // @Prod, block here

	// modules
	// heartbeat signal

	// receive message,

	// metrics, send to backend

	// receiv message and hanle and sendback

	fmt.Println("Agent start func - end")

	return nil
}

func (agent *Agent) RunPeriodicTask() {
	fmt.Println("SetUp for periodic task - start")

	interval := 5 * time.Second
	ticker := time.NewTicker(interval)

	for range ticker.C {
		fmt.Println("RunPeriodicTask - start")

		responseData, err := requestWithBackend.GetOperationCommandFromBackend()
		if err != nil {
			continue
		}
		fmt.Println("responseData in RunPeriodicTask\n", responseData)
		operationScript := responseData.Result.OperationScript
		fmt.Println("our operationScript:", operationScript)
		stdOut, err := runCommand.RunCommandByScriptContent(operationScript)
		var returnError bool = false
		var stdErr string = ""
		if err != nil {
			stdErr = stdOut + "\n======\n" + err.Error()
			stdOut = ""
			returnError = true
		}
		fmt.Println("stdOut:", stdOut)
		fmt.Println("stdErr:", stdErr)

		postResult := TRequest.OneOperationCommand{
			Id:               responseData.Result.Id,
			OperationCommand: responseData.Result.OperationCommand,
			Status:           responseData.Result.Status,
			OperationScript:  responseData.Result.OperationScript,
			OperationResult: TRequest.OperationResult{
				StdOut: stdOut,
				// StdErr:     stdErr.Error(),
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
