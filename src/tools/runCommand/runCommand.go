package runCommand

import (
	"fmt"
	"os/exec"
	"runtime"
)

const (
	Linux = iota
	Windows
	Unknown
)

func getOSName() int {
	os := runtime.GOOS
	if os == "windows" {
		fmt.Println("Windows operating system detected")
		return Windows
	} else if os == "linux" {
		fmt.Println("Linux operating system detected")
		return Linux
	}
	fmt.Printf("Unknown operating system: %s\n", os)
	return Unknown
}

func RunCommand() []byte {

	OSNameEnum := getOSName()

	if OSNameEnum == Linux {
		fmt.Println("package runCommand - RunCommand - Linux")
		cmd := exec.Command("sh", "-c", "/Users/jieanyang/Documents/freelancer_work/ansys/HelloWorldGoAgent/src/common/runCommand/script.sh")
		output, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Printf("Error running script: %s\n", err)
		} else {
			fmt.Printf("Script output: %s\n", output)
		}

		return output

	}

	if OSNameEnum == Windows {
		fmt.Println("package runCommand - RunCommand - Windows")
		// cmd := exec.Command("cmd", "/C", "C:\\Users\\jieanyang\\Documents\\freelancer_work\\ansys\\HelloWorldGoAgent\\src\\common\\runCommand\\script.bat")
		cmd := exec.Command("powershell.exe", "/Users/jieanyang/Documents/freelancer_work/ansys/HelloWorldGoAgent/src/common/runCommand/2023-02-27-first_script.ps1")
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

func RunCommandByScriptContent(scriptContent string) []byte {
	fmt.Println("package runCommand - RunCommandByContent")

	cmd := exec.Command("sh", "-c", scriptContent)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Error running script: %s\n", err)
	} else {
		fmt.Printf("Script output: %s\n", output)
	}

	return output
}
