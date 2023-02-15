package agentHttpController

import (
	"fmt"
	"net/http"
)

func HomeCtroller(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Go - Hello World</h1>")
}
