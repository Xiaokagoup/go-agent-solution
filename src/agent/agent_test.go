package agent

import "testing"

func TestAgent(t *testing.T) {
	agent := NewAgent()
	agent.Start()
}
