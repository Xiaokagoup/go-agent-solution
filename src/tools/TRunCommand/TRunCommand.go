package TRunCommand

import (
	"AnsysCSPAgent/src/tools/4_base/TOS"
	"AnsysCSPAgent/src/tools/4_base/TRequest"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

// === runCommand with url - start ===
type RequestDataForRunCommandByScriptContent struct {
	ScriptContent string `json:"scriptContent" default:"#!/bin/bash\necho \"start\"\necho \"hello yang\"\necho \"end\""`
	// ScriptContent string `json:"scriptContent" default:"Write-Output \"Windos PowerShell\"\nWrite-Output \"start\"\nWrite-Output \"hello yang\"\nWrite-Output \"end\""`
}
type RequestDataForRunCommandByUrl struct {
	Url string `json:"url" default:"https://ansys-gateway-development.s3.eu-west-3.amazonaws.com/2023-02-27-linux-script.sh"`
	// Url string `json:"url" default:"https://ansys-gateway-development.s3.eu-west-3.amazonaws.com/2023-02-27-windows-script.ps1"`
}

type RunCommandWithUrlRequestData struct {
	Value string `json:"value"`
	RequestDataForRunCommandByUrl
	RequestDataForRunCommandByScriptContent
}

func RunCommandByUrl(url string) (string, error) {

	// Get content from url
	responseFromScriptUrl, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer responseFromScriptUrl.Body.Close()
	fmt.Println("responseFromScriptUrl", responseFromScriptUrl)
	scriptContent, err := ioutil.ReadAll(responseFromScriptUrl.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("string(scriptContent)", string(scriptContent))

	scriptOutput, _ := RunCommandByScriptContent(string(scriptContent))
	fmt.Println("scriptOutput", scriptOutput)

	data := TRequest.StringRequestData{Result: string(scriptOutput)}

	return data.Result, nil
}

// === runCommand with url - end ===

// === runCommand with scriptContent - start ===
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

// === runCommand with scriptContent - end ===
