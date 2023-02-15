package agent

import (
	"errors"
	"fmt"
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

func (agent *Agent) Start() error {
	if agent.state != Waiting {
		return WrongStateError
	}

	agent.state = Running
	fmt.Println("Start - agent", agent.state)

	for {
	}

	return nil
}
