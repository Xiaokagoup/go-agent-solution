package agent

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"

	agtHttp "github.com/JieanYang/HelloWorldGoAgent/src/agentHttp"
	"github.com/JieanYang/HelloWorldGoAgent/src/tools/agentMetadataManager"
)

const (
	Waiting = iota
	Running
)

var WrongStateError = errors.New("Can't take the operation in the current state")

var BACKEND_ENDPOINT string = "https://13b1-2a01-cb16-60-e0c3-1440-9ccb-d2a9-6d4e.ngrok-free.app"

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

func RunPeriodicTask() {
	interval := 5 * time.Second
	ticker := time.NewTicker(interval)

	for range ticker.C {
		SendPOSTRequest(BACKEND_ENDPOINT+"/node/api-docs", runtime.GOOS+" | I'm a request in RunPeriodicTask")
	}
}

func SendPOSTRequest(url string, messages string) *http.Response {
	type POSTData struct {
		Messages string `json:"message"`
	}

	postData := &POSTData{
		Messages: messages,
	}
	jsonPostData, err := json.Marshal(postData)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPostData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return resp
}

func (agent *Agent) Init() {
	// load config
	// load modules
	// load heartbeat signal
	// load metrics
	// load message

	RunPeriodicTask()

	psk := GeneratePSK_key()
	agentMetadataManager.GetOrCreateConfigFileWithSpecifiedPskKey(psk) // save psk key to config file

	resp := SendPOSTRequest(BACKEND_ENDPOINT+"/node/api-docs", runtime.GOOS+" | Hello World")

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status code: %d", resp.StatusCode))
	}

	// Print response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println("Response body:", bodyString)
}
