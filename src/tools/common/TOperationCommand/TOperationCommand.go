//
// Copyright (C) 2023 ANSYS, Inc. Unauthorized use, distribution, or duplication is prohibited.
//

package TOperationCommand

import "fmt"

type OperationResult struct {
	StdOut      string `json:"stdOut"`
	StdErr      string `json:"stdErr"`
	ReturnError bool   `json:"returnError"`
}

type OneOperationCommand struct {
	Id               string          `json:"id"`
	OperationCommand string          `json:"operationCommand"`
	Status           string          `json:"status"`
	OperationScript  string          `json:"operationScript"`
	OperationResult  OperationResult `json:"operationResult"`
	TryTimes         int             `json:"tryTimes"`
}

type OneOperationCommandResponseData struct {
	Result OneOperationCommand `json:"result"`
}

func (rd OneOperationCommandResponseData) String() string {
	return fmt.Sprintf("===============\nID: %s\nOperation Command: %s\nStatus: %s\nOperation Script: %s\nStdOut: %s\nStdErr: %s\nReturn Error: %t\nTry Times: %d\n===============\n",
		rd.Result.Id,
		rd.Result.OperationCommand,
		rd.Result.Status,
		rd.Result.OperationScript,
		rd.Result.OperationResult.StdOut,
		rd.Result.OperationResult.StdErr,
		rd.Result.OperationResult.ReturnError,
		rd.Result.TryTimes)
}
