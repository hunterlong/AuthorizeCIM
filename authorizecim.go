package AuthorizeCIM

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

var api_endpoint string = "https://apitest.authorize.net/xml/v1/request.api"
var apiName string
var apiKey string
var testMode string

//
//var CurrentUser User

func SetAPIInfo(name string, key string, mode string) {
	apiKey = key
	apiName = name
	if mode == "test" {
		testMode = "testMode"
		api_endpoint = "https://apitest.authorize.net/xml/v1/request.api"
	} else {
		testMode = "liveMode"
		api_endpoint = "https://api.authorize.net/xml/v1/request.api"
	}
}

//func MakeUser(userID string) User {
//	CurrentUser = User{ID: "55", Email: userID, ProfileID: "0"}
//	return CurrentUser
//}

func SendRequest(input []byte) ([]byte) {
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
	return body
}

//func TestConnection() bool {
//	authToken := AuthenticateTestRequest{MerchantAuthentication{Name: apiName, TransactionKey: apiKey}}
//	authnettest := AuthorizeNetTest{AuthenticateTestRequest: authToken}
//	jsoned, _ := json.Marshal(authnettest)
//	outgoing, _ := SendRequest(string(jsoned))
//	status, _ := FindResultCode(outgoing)
//	return status
//}

//func FindResultCode(incoming map[string]interface{}) (bool, string) {
//	messages, _ := incoming["messages"].(map[string]interface{})
//
//	if messages != nil {
//
//		if messages["resultCode"].(string) == "Ok" {
//			return true, ""
//		} else {
//			messagesInfo := messages["message"].(map[string]interface{})
//			return false, messagesInfo["text"].(string)
//		}
//
//	}
//
//	return false, ""
//}
//
//func UpdateSubscription() {
//
//}

func GetSubscriptions() {

}
