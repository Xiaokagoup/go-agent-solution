package runCommand

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("sh", "-c", "./script.sh")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Error running script: %s\n", err)
	} else {
		fmt.Printf("Script output: %s\n", output)
	}
}
