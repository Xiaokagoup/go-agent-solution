package runCommand

import (
	"fmt"
	"testing"

	"github.com/JieanYang/HelloWorldGoAgent/src/common/logger"
)

func TestRunCommand(t *testing.T) {
	fmt.Println("Start test - runCommand.go")
	logger.Log(("yang"))

	runCommand()

	fmt.Println("End test - runCommand.go")
}
