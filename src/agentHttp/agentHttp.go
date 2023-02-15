package agentHttp

import (
	"fmt"
	"net/http"
	"sync"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Go - Hello World</h1>")
}
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
