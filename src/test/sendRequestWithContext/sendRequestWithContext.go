package sendRequestWithContext

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func requestData(url string) (*http.Response, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}

func main() {
	fmt.Println("Hello World!")

	interval := 5 * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context timed out, stopping.")
			return
		case <-ticker.C:
			fmt.Println("Call - start")

			url := "https://ff66-2a01-cb16-60-e0c3-a023-d11-c4af-ce6a.ngrok-free.app/node/aws/receiveOperationCommandResult"
			response, err := requestData(url)

			if err != nil {
				var netErr *net.OpError
				if errors.As(err, &netErr) && netErr.Timeout() {
					fmt.Println("Request timed out")
				} else {
					fmt.Println("Error:", err)
				}
			} else {
				body, _ := ioutil.ReadAll(response.Body)
				fmt.Println("Response Status:", response.Status)
				fmt.Println("Response Body:", string(body))
			}

			fmt.Println("Call - end")
		}
	}
}
