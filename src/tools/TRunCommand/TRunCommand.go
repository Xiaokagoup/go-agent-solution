package TRunCommand

import (
	"AnsysCSPAgent/src/tools/4_base/TOS"
	"errors"
	"fmt"
	"os/exec"
)

func RunCommandByScriptContent(scriptContent string) (string, error) {

	OSNameEnum := TOS.GetOSName()

	fmt.Println("scriptContent", scriptContent)

	if OSNameEnum == TOS.Linux || OSNameEnum == TOS.MacOS {
		fmt.Println("package runCommand - RunCommandByContent - Linux")

		cmd := exec.Command("sh", "-c", scriptContent)
		output, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Printf("Error running script: %s\n", err)
		} else {
			fmt.Printf("Script output: %s\n", output)
		}

		return string(output), err
	}

	if OSNameEnum == TOS.Windows {
		fmt.Println("package runCommand - RunCommand - Windows")
		// cmd := exec.Command("cmd", "/C", "C:\\Users\\jieanyang\\Documents\\freelancer_work\\ansys\\HelloWorldGoAgent\\src\\common\\runCommand\\script.bat")
		cmd := exec.Command("powershell.exe", scriptContent)
		output, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Printf("Error running script: %s\n", err)
		} else {
			fmt.Printf("Script output: %s\n", output)
		}

		return string(output), err
	}

	return "", errors.New("the system OS is not supported")
}
