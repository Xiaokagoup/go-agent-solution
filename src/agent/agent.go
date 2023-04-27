package agent

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	agtHttp "github.com/JieanYang/HelloWorldGoAgent/src/agentHttp"
	"github.com/JieanYang/HelloWorldGoAgent/src/tools/agentMetadataManager"
)

const (
	Waiting = iota
	Running
)

var WrongStateError = errors.New("Can't take the operation in the current state")

var BACKEND_ENDPOINT string = "https://ff66-2a01-cb16-60-e0c3-a023-d11-c4af-ce6a.ngrok-free.app"

type Agent struct {
	state int
}

func NewAgent() *Agent {
	agent := Agent{
		state: Waiting,
	}

	return &agent
}

// Business logic
func (agent *Agent) Start() error {
	if agent.state != Waiting {
		return WrongStateError
	}

	agent.state = Running
	fmt.Println("Start - agent", agent.state)

	agent.Init()
	agtHttp.StartHttp()
	// modules
	// heartbeat signal

	// receive message,

	// metrics, send to backend

	// receiv message and hanle and sendback

	return nil
}

func GeneratePSK_key() string {
	// Generate PSK key
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	psk := base64.StdEncoding.EncodeToString(key)
	return psk
}

func GetOperationCommandFromBackend() {

}

func RunPeriodicTask() {
	interval := 5 * time.Second
	ticker := time.NewTicker(interval)

	fmt.Println("RunPeriodicTask - start")
	for range ticker.C {
		responseData, err := SendGETRequest(BACKEND_ENDPOINT + "/node/aws/getMockOperationCommand")
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("%+v\n", responseData)
		}
	}
	fmt.Println("RunPeriodicTask - end")
}

type OperationResult struct {
	ReturnCode int    `json:"returnCode"`
	StdOut     string `json:"stdOut"`
	StdErr     string `json:"stdErr"`
}
type OneOperationCommand struct {
	Id               string          `json:"id"`
	OperationCommand string          `json:"operationCommand"`
	Status           string          `json:"status"`
	OperationScript  string          `json:"operationScript"`
	OperationResult  OperationResult `json:"operationResult"`
	TryTimes         int             `json:"tryTimes"`
}
type ResponseData struct {
	Result OneOperationCommand `json:"result"`
}

func (rd ResponseData) String() string {
	return fmt.Sprintf("ID: %s\nOperation Command: %s\nStatus: %s\nOperation Script: %s\nReturn Code: %d\nStdOut: %s\nStdErr: %s\nTry Times: %d",
		rd.Result.Id,
		rd.Result.OperationCommand,
		rd.Result.Status,
		rd.Result.OperationScript,
		rd.Result.OperationResult.ReturnCode,
		rd.Result.OperationResult.StdOut,
		rd.Result.OperationResult.StdErr,
		rd.Result.TryTimes)
}

func SendGETRequest(url string) (*ResponseData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseData ResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

// func SendPOSTRequest(url string, messages string) ResponseData {
// 	type POSTData struct {
// 		Messages string `json:"message"`
// 	}

// 	postData := &POSTData{
// 		Messages: messages,
// 	}
// 	jsonPostData, err := json.Marshal(postData)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPostData))
// 	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonPostData))
// 	if err != nil {
// 		panic(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	fmt.Println("Response body:", resp)
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var data ResponseData
// 	err = json.Unmarshal(body, &data)
// 	if err != nil {
// 		// handle error
// 	}

// 	return data
// }

func (agent *Agent) Init() {
	// load config
	// load modules
	// load heartbeat signal
	// load metrics
	// load message

	RunPeriodicTask()

	psk := GeneratePSK_key()
	agentMetadataManager.GetOrCreateConfigFileWithSpecifiedPskKey(psk) // save psk key to config file
}
