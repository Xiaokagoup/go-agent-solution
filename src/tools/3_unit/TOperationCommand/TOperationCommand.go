package TOperationCommand

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
