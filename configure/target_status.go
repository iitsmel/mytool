package configure

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const timeout = 5 * time.Second

func MakeRequest(url string) (*http.Response, error) {
	client := &http.Client{Timeout: timeout}
	return client.Get(url)
}

func ShowResponse(passing *Scope) {
	URL := passing.Attack.Target
	responsing, errormsg := MakeRequest(URL)

	if errormsg != nil {
		fmt.Printf("Error making http request: %s\n", errormsg)
		return
	}

	fmt.Println("Status Code      :  ", responsing.StatusCode)
	fmt.Println("Protocal Version :  ", responsing.Proto)
	fmt.Println("Response Headers :  ", responsing.Header)
	fmt.Println("Content  Length  :  ", responsing.ContentLength)
	fmt.Println("Transfer Encoding:  ", responsing.TransferEncoding)
}

func CheckTargetStatus(passing *Scope) bool {
	response, errorStatus := http.Get(passing.Attack.Target)

	if errorStatus != nil {
		fmt.Printf("HTTP request failed: %v\n", errorStatus)
		os.Exit(0)
	}

	defer response.Body.Close()

	return response.StatusCode != http.StatusOK
}
