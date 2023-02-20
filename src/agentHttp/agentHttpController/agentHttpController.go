package agentHttpController

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/JieanYang/HelloWorldGoAgent/src/tools/runCommand"
	"github.com/gin-gonic/gin"
)

type Reponse struct {
	Results string
}

type RequestData struct {
	Value         string `json:"value"`
	Url           string `json:"url"`
	ScriptContent string `json:"scriptContent"`
}

func HomeGetController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func HomePostController(c *gin.Context) {
	var reqData RequestData
	if err := c.ShouldBindJSON(&reqData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// fmt.Println("reqData", reqData)

	c.JSON(http.StatusOK, gin.H{"results": reqData})
}

// func RunCommandWithFormData(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		err := r.ParseForm()
// 		if err != nil {
// 			http.Error(w, "Error parsing form data", http.StatusBadRequest)
// 			return
// 		}
// 		name := r.FormValue("name")

// 		fmt.Fprintf(w, "<h1>Paris Hello, %s</h1>", name)
// 	} else {
// 		fmt.Fprint(w, "<h1>Hellow Word</h1>")
// 	}

// }

// func RunCommandWithBody(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		var body struct{ Name string }
// 		err := json.NewDecoder(r.Body).Decode(&body)
// 		if err != nil {
// 			http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
// 			return
// 		}

// 		output := runCommand.RunCommand()

// 		resultsObj := Reponse{Results: string(output)}
// 		data, err := json.Marshal(resultsObj)
// 		if err != nil {
// 			http.Error(w, "Error generate JSON reuslts", http.StatusBadRequest)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(data)

// 		// fmt.Fprintf(w, "<h1>Hello, %s</h1>", body.Name)
// 	} else {
// 		fmt.Fprint(w, "<h1>Hellow Word</h1>")
// 	}

// }

func RunCommandByScriptContent(c *gin.Context) {
	var reqData RequestData
	if err := c.ShouldBindJSON(&reqData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	scriptOutput := runCommand.RunCommandByScriptContent(string(reqData.ScriptContent))

	data := Reponse{Results: string(scriptOutput)}

	c.JSON(http.StatusOK, gin.H{"results": data})

}

func RunCommandByUrl(c *gin.Context) {
	var reqData RequestData
	if err := c.ShouldBindJSON(&reqData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("reqData", reqData)
	fmt.Println("reqData.Url", reqData.Url)

	// Get content from url
	responseFromScriptUrl, err := http.Get(reqData.Url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer responseFromScriptUrl.Body.Close()
	fmt.Println("responseFromScriptUrl", responseFromScriptUrl)
	scriptContent, err := ioutil.ReadAll(responseFromScriptUrl.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("string(scriptContent)", string(scriptContent))

	scriptOutput := runCommand.RunCommandByScriptContent(string(scriptContent))
	fmt.Println("scriptOutput", scriptOutput)

	data := Reponse{Results: string(scriptOutput)}

	c.JSON(http.StatusOK, gin.H{"results": data})

}
