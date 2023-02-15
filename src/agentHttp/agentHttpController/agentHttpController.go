package agentHttpController

import (
	"fmt"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Go - Hello World</h1>")
}

func RunCommandControllerWithFormData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")

		fmt.Fprintf(w, "<h1>Hello, %s</h1>", name)
	} else {
		fmt.Fprint(w, "<h1>Hellow Word</h1>")
	}

}
