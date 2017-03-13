package AuthorizeCIM

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

var api_endpoint string = "https://apitest.authorize.net/xml/v1/request.api"
var apiName *string
var apiKey *string
var testMode string
var showLogs bool = true

func SetAPIInfo(name string, key string, mode string) {
	apiKey = &key
	apiName = &name
	if mode == "live" {
		showLogs = false
		testMode = "liveMode"
		api_endpoint = "https://api.authorize.net/xml/v1/request.api"
	} else {
		showLogs = true
		testMode = "testMode"
		api_endpoint = "https://apitest.authorize.net/xml/v1/request.api"
	}
}

func IsConnected() bool {
	info := GetMerchantDetails()
	if info.Ok() {
		return true
	}
	return false
}

func GetAuthentication() MerchantAuthentication {
	auth := MerchantAuthentication{
		Name:           apiName,
		TransactionKey: apiKey,
	}
	return auth
}

func SendRequest(input []byte) []byte {
	req, err := http.NewRequest("POST", api_endpoint, bytes.NewBuffer(input))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	if showLogs {
		fmt.Println(string(body))
	}
	return body
}

func (r AVS) Text() string {
	var response string
	switch r.avsResultCode {
	case "E":
		response = "AVS data provided is invalid or AVS is not allowed for the card type that was used."
	case "R":
		response = "The AVS system was unavailable at the time of processing."
	case "G":
		response = "The card issuing bank is of non-U.S. origin and does not support AVS"
	case "U":
		response = "The address information for the cardholder is unavailable."
	case "S":
		response = "The U.S. card issuing bank does not support AVS."
	case "N":
		response = "Address: No Match ZIP Code: No Match"
	case "A":
		response = "Address: Match ZIP Code: No Match"
	case "Z":
		response = "Address: No Match ZIP Code: Match"
	case "W":
		response = "Address: No Match ZIP Code: Matched 9 digits"
	case "X":
		response = "Address: Match ZIP Code: Matched 9 digits"
	case "Y":
		response = "Address: Match ZIP: Matched first 5 digits"
	}
	return response
}
