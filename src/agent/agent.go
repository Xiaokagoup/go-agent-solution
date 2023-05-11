package agent

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	agtHttp "github.com/JieanYang/HelloWorldGoAgent/src/agentHttp"
	"github.com/JieanYang/HelloWorldGoAgent/src/tools/agentMetadataManager"
	"github.com/JieanYang/HelloWorldGoAgent/src/tools/requestWithBackend"
	"github.com/JieanYang/HelloWorldGoAgent/src/tools/runCommand"
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

// Business logic
func (agent *Agent) Start() error {
	fmt.Println("Agent start func - start")

	if agent.state != Waiting {
		return ErrWrongState
	}

	agent.state = Running
	fmt.Println("Start - agent", agent.state)

	agent.Init()
	go agtHttp.StartHttp()
	agent.RunPeriodicTask() // block here

	// modules
	// heartbeat signal

	// receive message,

	// metrics, send to backend

	// receiv message and hanle and sendback

	fmt.Println("Agent start func - end")

	return nil
}

func GeneratePSK_key() string {
	// Generate PSK key
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	psk := base64.StdEncoding.EncodeToString(key)
	return psk
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

		postResult := requestWithBackend.OneOperationCommand{
			Id:               responseData.Result.Id,
			OperationCommand: responseData.Result.OperationCommand,
			Status:           responseData.Result.Status,
			OperationScript:  responseData.Result.OperationScript,
			OperationResult: requestWithBackend.OperationResult{
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

func (agent *Agent) Init() {
	// load config
	// load modules
	// load heartbeat signal
	// load metrics
	// load message

	psk := GeneratePSK_key()
	agentMetadataManager.GetOrCreateConfigFileWithSpecifiedPskKey(psk) // save psk key to config file

}
