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

var WrongStateError = errors.New("Can't take the operation in the current state")

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
	if agent.state != Waiting {
		return WrongStateError
	}

	agent.state = Running
	fmt.Println("Start - agent", agent.state)

	agent.Init()
	agtHttp.StartHttp()
	// modules
	// heartbeat signal

	// receive message,

	// metrics, send to backend

	// receiv message and hanle and sendback

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

func RunPeriodicTask() {
	interval := 5 * time.Second
	ticker := time.NewTicker(interval)

	fmt.Println("RunPeriodicTask - start")
	for range ticker.C {
		responseData := requestWithBackend.GetOperationCommandFromBackend()
		fmt.Println("responseData in RunPeriodicTask\n", responseData)
		operationScript := responseData.Result.OperationScript
		fmt.Println("our operationScript:", operationScript)
		stdOut, stdErr := runCommand.RunCommandByScriptContent(operationScript)
		fmt.Println("stdOut:", stdOut)
		fmt.Println("stdErr:", stdErr)

		postResult := requestWithBackend.OneOperationCommand{
			Id:               responseData.Result.Id,
			OperationCommand: responseData.Result.OperationCommand,
			Status:           responseData.Result.Status,
			OperationScript:  responseData.Result.OperationScript,
			OperationResult: requestWithBackend.OperationResult{
				ReturnCode: 200,
				StdOut:     stdOut,
				// StdErr:     stdErr.Error(),
				StdErr: "",
			},
			TryTimes: responseData.Result.TryTimes,
		}

		requestWithBackend.PostOperationCommandResultToBackend(postResult)
	}
	fmt.Println("RunPeriodicTask - end")
}

func (agent *Agent) Init() {
	// load config
	// load modules
	// load heartbeat signal
	// load metrics
	// load message

	RunPeriodicTask()

	psk := GeneratePSK_key()
	agentMetadataManager.GetOrCreateConfigFileWithSpecifiedPskKey(psk) // save psk key to config file
}
