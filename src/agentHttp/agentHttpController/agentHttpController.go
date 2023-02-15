package agentHttpController

import (
	"fmt"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Go - Hello World</h1>")
}

func RunCommandController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Go - Run command</h1>")
}
