package agentHttp

import (
	"net/http"
	"sync"

	agentHttpController "github.com/JieanYang/HelloWorldGoAgent/src/agentHttp/agentHttpController"
)

func StartHttp() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		http.HandleFunc("/", agentHttpController.HomeController)
		http.HandleFunc("/runCommandWithFormData", agentHttpController.RunCommandWithFormData)
		http.HandleFunc("/runCommandWithBody", agentHttpController.RunCommandWithBody)

		err := http.ListenAndServe(":9001", nil) // Block code
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()

	wg.Wait()

}
