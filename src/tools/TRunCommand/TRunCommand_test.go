//
// Copyright (C) 2023 ANSYS, Inc. Unauthorized use, distribution, or duplication is prohibited.
//

package TRunCommand

import (
	"fmt"
	"os/exec"
	"testing"

	"AnsysCSPAgent/src/tools/4_base/TOS"
)

func RunCommandTest() []byte {

	OSNameEnum := TOS.GetOSName()

	if OSNameEnum == TOS.Linux || OSNameEnum == TOS.MacOS {
		fmt.Println("package runCommand - RunCommand - Linux")
		cmd := exec.Command("sh", "-c", "/Users/jieanyang/Documents/freelancer_work/ansys/ansysCSPAgent/src/common/runCommand/script.sh")
		output, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Printf("Error running script: %s\n", err)
		} else {
			fmt.Printf("Script output: %s\n", output)
		}

		return output

	}

	if OSNameEnum == TOS.Windows {
		fmt.Println("package runCommand - RunCommand - Windows")
		// cmd := exec.Command("cmd", "/C", "C:\\Users\\jieanyang\\Documents\\freelancer_work\\ansys\\ansysCSPAgent\\src\\common\\runCommand\\script.bat")
		cmd := exec.Command("powershell.exe", "/Users/jieanyang/Documents/freelancer_work/ansys/ansysCSPAgent/src/common/runCommand/2023-02-27-first_script.ps1")
		output, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Printf("Error running script: %s\n", err)
		} else {
			fmt.Printf("Script output: %s\n", output)
		}

		return output
	}

	return nil
}

func TestRunCommand(t *testing.T) {
	fmt.Println("Start test - runCommand.go")

	RunCommandTest()

	fmt.Println("End test - runCommand.go")
}
