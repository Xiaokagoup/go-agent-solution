package agent

import (
	"errors"
	"fmt"

	agtHttp "github.com/JieanYang/HelloWorldGoAgent/src/agentHttp"
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

	agtHttp.StartHttp()
	// modules
	// heartbeat signal

	// receive message,

	// metrics, send to backend

	// receiv message and hanle and sendback

	return nil
}
