package agentHttpController

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JieanYang/HelloWorldGoAgent/src/common/runCommand"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Go - Hello World</h1>")
}

func RunCommandWithFormData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")

		fmt.Fprintf(w, "<h1>Paris Hello, %s</h1>", name)
	} else {
		fmt.Fprint(w, "<h1>Hellow Word</h1>")
	}

}

func RunCommandWithBody(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var body struct{ Name string }
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
			return
		}

		output := runCommand.RunCommand()
		// outputString := string(output)

		w.Header().Set("Content-Type", "application/json")
		w.Write(output)

		// fmt.Fprintf(w, "<h1>Hello, %s</h1>", body.Name)
	} else {
		fmt.Fprint(w, "<h1>Hellow Word</h1>")
	}

}
