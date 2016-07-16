package AuthorizeCIM

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

var api_endpoint string
var apiName string
var apiKey string
var testMode string

var CurrentUser User

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

func MakeUser (userID string) User {
	CurrentUser = User{ID: "55", Email: userID, ProfileID: "0"}
	return CurrentUser
}

func CreateCustomerProfile(userInfo AuthUser) (string, bool) {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	profile := Profile{MerchantCustomerID: userInfo.Uuid, Description: userInfo.Description, Email: userInfo.Email}
	request := CreateCustomerProfileRequest{authToken, profile}
	newprofile := NewCustomerProfile{request}
	jsoned, _ := json.Marshal(newprofile)
	outgoing, _ := SendRequest(string(jsoned))
	success := FindResultCode(outgoing)
	var new_uuid string
	if success {
		new_uuid = outgoing["customerProfileId"].(string)
		CurrentUser.ProfileID = new_uuid
	} else {
		new_uuid = "0"
	}
	return new_uuid, success
}


func GetCustomerProfile(profileID string) (map[string]interface{}, bool) {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	profile := getCustomerProfileRequest{authToken, profileID}
	input := CustomerProfile{profile}
	jsoned, _ := json.Marshal(input)
	outgoing, _ :=SendRequest(string(jsoned))
	success := FindResultCode(outgoing)
	fmt.Println(outgoing)
	userProfile := outgoing["profile"].(map[string]interface{})
	fmt.Println(userProfile)
	return outgoing, success
}


func GetAllProfiles() []interface{} {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	profilerequest := getCustomerProfileIdsRequest{authToken}
	all := AllCustomerProfileIds{profilerequest}
	jsoned, _ := json.Marshal(all)
	outgoing, _ :=SendRequest(string(jsoned))
	return outgoing["ids"].([]interface{})
}


func DeleteCustomerProfile(profileID string) bool {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	profile := deleteCustomerProfileRequest{authToken, profileID}
	input := deleteCustomerProfile{profile}
	jsoned, _ := json.Marshal(input)
	outgoing, _ := SendRequest(string(jsoned))
	status := FindResultCode(outgoing)
	return status
}


func CreateCustomerBillingProfile(profileID string, creditCard CreditCard, address Address) (string, bool) {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	paymentProfile := PaymentBillingProfile{Address: address, Payment: Payment{CreditCard:creditCard}}
	request := CreateCustomerBillingProfileRequest{authToken, profileID, paymentProfile, testMode}
	newprofile := NewCustomerBillingProfile{request}
	jsoned, _ := json.Marshal(newprofile)
	outgoing, _ :=SendRequest(string(jsoned))
	status := FindResultCode(outgoing)
	var new_paymentID string
	if status {
		new_paymentID = outgoing["customerPaymentProfileId"].(string)
	} else {
		new_paymentID = "0"
	}
	return new_paymentID, status
}



func GetCustomerPaymentProfile(profileID string, paymentID string) (map[string]interface{}, bool) {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	profile := CustomerPaymentProfileRequest{authToken, profileID, paymentID}
	input := getCustomerPaymentProfileRequest{profile}
	jsoned, _ := json.Marshal(input)
	outgoing, _ := SendRequest(string(jsoned))
	success := FindResultCode(outgoing)
	fmt.Println(CurrentUser)
	return outgoing["paymentProfile"].(map[string]interface{}), success
}


func UpdateCustomerPaymentProfile(profileID string, paymentID string, creditCard CreditCard, address Address) bool {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	new_billing := UpdatePaymentBillingProfile{Address: address, Payment: Payment{CreditCard:creditCard}, CustomerPaymentProfileId: paymentID}
	profile := updateCustomerPaymentProfileRequest{authToken, profileID, new_billing, testMode}
	input := changeCustomerPaymentProfileRequest{profile}
	jsoned, _ := json.Marshal(input)
	outgoing, _ := SendRequest(string(jsoned))
	status := FindResultCode(outgoing)
	return status
}


func DeleteCustomerPaymentProfile(profileID string, paymentID string) bool {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	profile := deleteCustomerPaymentProfile{authToken, profileID, paymentID}
	input := deleteCustomerPaymentProfileRequest{profile}
	jsoned, _ := json.Marshal(input)
	outgoing, _ :=SendRequest(string(jsoned))
	status := FindResultCode(outgoing)
	return status
}

func SendRequest(input string) (map[string]interface{}, interface{}) {
	req, err := http.NewRequest("POST", api_endpoint, bytes.NewBuffer([]byte(input)))
	req.Header.Set("Content-Type", "application/json")
	errors := false
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	var dat map[string]interface{}
	//fmt.Printf(string(body))
	err = json.Unmarshal(body, &dat)
	if err!=nil {
		panic(err)
	}
	return dat, errors
}



func CreateTransaction(profileID string, paymentID string, item LineItem, amount string) (map[string]interface{}, bool, bool) {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	items := LineItems{LineItem: item}
	subProfile := SubProfile{CustomerPaymentProfileId: paymentID}
	transProfile := TranProfile{CustomerProfileId: profileID, SubProfile: subProfile}
	transaction := TransactionRequest{TransactionType: "authCaptureTransaction", Amount: amount, TranProfile: transProfile, LineItems: items}
	tranxrequest := CreateTransactionRequest{MerchantAuthentication: authToken, RefID: "none33", TransactionRequest: transaction}
	doTranx := DoCreateTransaction{tranxrequest}
	jsoned, _ := json.Marshal(doTranx)
	outgoing, _ := SendRequest(string(jsoned))
	var status, approved bool
	var response map[string]interface{}
	if outgoing["responseCode"]!=nil {
		if outgoing["responseCode"].(string) != "1" {
			approved = false
			status = true
			response = map[string]interface{}{}
		} else {
			status = FindResultCode(outgoing)
			approved = TransactionApproved(outgoing)
			response = outgoing["transactionResponse"].(map[string]interface{})
		}
	} else {
		approved = false
		status = false
		response = map[string]interface{}{}
	}

	return response, approved, status
}


func TestConnection() bool {
	authToken := AuthenticateTestRequest{MerchantAuthentication{Name: apiName, TransactionKey: apiKey}}
	authnettest := AuthorizeNetTest{AuthenticateTestRequest:authToken}
	jsoned, _ := json.Marshal(authnettest)
	outgoing, _ := SendRequest(string(jsoned))
	status := FindResultCode(outgoing)
	return status
}


func CreateShippingAddress(profileID string, address Address) (string, bool) {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	customerShipping := CustomerShippingAddress{authToken,profileID,address}
	customerShippingRequest := CustomerShippingAddressRequest{customerShipping}
	jsoned, _ := json.Marshal(customerShippingRequest)
	outgoing, _ := SendRequest(string(jsoned))
	success := FindResultCode(outgoing)
	var new_address_id string
	if !success {
		new_address_id = "0"
	} else {
		new_address_id = outgoing["customerAddressId"].(string)
	}
	return new_address_id, success
}

func GetShippingAddress(profileID string, shippingID string) (map[string]interface{}, bool) {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	customerShipping := GetCustomerShippingAddress{authToken,profileID,shippingID}
	customerShippingRequest := GetCustomerShippingAddressRequest{customerShipping}
	jsoned, _ := json.Marshal(customerShippingRequest)
	outgoing, _ := SendRequest(string(jsoned))
	success := FindResultCode(outgoing)
	return outgoing["address"].(map[string]interface{}), success
}

func DeleteShippingAddress(profileID string, shippingID string) bool {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	customerShipping := GetCustomerShippingAddress{authToken,profileID,shippingID}
	customerShippingRequest := DeleteCustomerShippingAddressRequest{customerShipping}
	jsoned, _ := json.Marshal(customerShippingRequest)
	outgoing, _ := SendRequest(string(jsoned))
	status := FindResultCode(outgoing)
	return status
}


func GetTransactionDetails(tranID string) map[string]interface{} {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	transDetails := TransactionDetails{authToken,tranID}
	transactionRequest := TransactionDetailsRequest{transDetails}
	jsoned, _ := json.Marshal(transactionRequest)
	outgoing, _ := SendRequest(string(jsoned))
	if outgoing["transaction"]!=nil {
		return outgoing["transaction"].(map[string]interface{})
	}
	return map[string]interface{}{}
}


func FindResultCode(incoming map[string]interface{}) bool {
	messages, _ := incoming["messages"].(map[string]interface{})
	if messages!=nil {
		if messages["resultCode"] == "Ok" {
			return true
		}
	}
	return false
}

func TransactionApproved(incoming map[string]interface{}) bool {
	if incoming!=nil {
		messages, _ := incoming["transactionResponse"].(map[string]interface{})
		if messages["responseCode"] == "1" {
			return true
		}
	}
	return false
}


func CreateSubscription(newSubscription Subscription) (string, bool) {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	subscriptonSubmit := CreateSubscriptionRequest{ARBCreateSubscription{authToken, newSubscription}}
	jsoned, _ := json.Marshal(subscriptonSubmit)
	outgoing, _ := SendRequest(string(jsoned))
	status := FindResultCode(outgoing)
	if status {
		return outgoing["subscriptionId"].(string), status
	}
	return "0", status
}



func RefundTransactions(){

}

func VoidTransaction(){

}

func DeleteSubscription(){

}

func UpdateSubscription(){

}

func GetSubscriptions(){

}

