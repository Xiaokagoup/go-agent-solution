package agentHttp

import (
	"net/http"
	"sync"
)

func StartHttp() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		http.HandleFunc("/", rootHandler)

		err := http.ListenAndServe(":9001", nil) // Block code
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()

	wg.Wait()

}
