package configure

import (
	"fmt"
	"net/http"
	"time"
)

const timeout = 5 * time.Second

func MakeRequest(url string) (*http.Response, error) {
	client := &http.Client{Timeout: timeout}
	return client.Get(url)
}

func ShowResponse(passing *Scope) {
	URL := passing.HTTP.URL
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

func CheckTargetStatus(target string) bool {
	response, errorStatus := http.Get(target)

	if errorStatus != nil || response.StatusCode != http.StatusOK {
		return false
	}

	defer response.Body.Close()
	return true
}
