package runCommand

import (
	"fmt"
	"testing"

	"github.com/JieanYang/HelloWorldGoAgent/src/logger"
)

func TestRunCommand(t *testing.T) {
	fmt.Println("Start test - runCommand.go")
	logger.Log(("yang"))

	main()

	fmt.Println("End test - runCommand.go")
}
