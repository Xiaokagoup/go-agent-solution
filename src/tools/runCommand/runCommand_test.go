package runCommand

import (
	"fmt"
	"testing"

	"AnsysCSPAgent/src/tools/logger"
)

func TestRunCommand(t *testing.T) {
	fmt.Println("Start test - runCommand.go")
	logger.Log(("yang"))

	RunCommandTest()

	fmt.Println("End test - runCommand.go")
}
