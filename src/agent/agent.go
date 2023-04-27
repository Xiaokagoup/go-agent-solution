package agent

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	agtHttp "github.com/JieanYang/HelloWorldGoAgent/src/agentHttp"
	"github.com/JieanYang/HelloWorldGoAgent/src/tools/agentMetadataManager"
	"github.com/JieanYang/HelloWorldGoAgent/src/tools/requestWithBackend"
)

const (
	Waiting = iota
	Running
)

var WrongStateError = errors.New("Can't take the operation in the current state")

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

	fmt.Println("RunPeriodicTask - start")
	for range ticker.C {
		requestWithBackend.GetOperationCommandFromBackend()
	}
	fmt.Println("RunPeriodicTask - end")
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
