package runCommand

import (
	"fmt"
	"os/exec"
)

func RunCommand() []byte {

	fmt.Println("RunCommand mobule")
	cmd := exec.Command("sh", "-c", "/Users/jieanyang/Documents/freelancer_work/ansys/HelloWorldGoAgent/src/common/runCommand/script.sh")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Error running script: %s\n", err)
	} else {
		fmt.Printf("Script output: %s\n", output)
	}

	return output
}
