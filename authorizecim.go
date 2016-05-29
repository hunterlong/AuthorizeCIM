package AuthorizeCIM

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

var api_endpoint string = "https://apitest.authorize.net/xml/v1/request.api"
var apiName string
var apiKey string


func SetAPIInfo(apiName string, apiKey string) {
	apiKey = apiKey
	apiName = apiName
}

func CreateCustomerProfile(userInfo AuthUser) (string, map[string]string) {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	profile := Profile{MerchantCustomerID: userInfo.Uuid, Description: userInfo.Description, Email: userInfo.Email}
	request := CreateCustomerProfileRequest{authToken, profile}
	newprofile := NewCustomerProfile{request}
	jsoned, _ := json.Marshal(newprofile)
	outgoing, _ := SendRequest(string(jsoned))
	var new_uuid string
	var errors map[string]string
	if outgoing["customerProfileId"]==nil {
		new_uuid = "0"
		errors = map[string]string{"message": "User cannot be created"}
	} else {
		new_uuid = outgoing["customerProfileId"].(string)
		errors = nil
	}
	return new_uuid, errors
}


func GetCustomerProfile(profileID string) map[string]interface{} {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	profile := getCustomerProfileRequest{authToken, profileID}
	input := CustomerProfile{profile}
	jsoned, _ := json.Marshal(input)
	outgoing, _ :=SendRequest(string(jsoned))
	return outgoing
}


func GetAllProfiles() map[string]interface{} {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	profilerequest := getCustomerProfileIdsRequest{authToken}
	all := AllCustomerProfileIds{profilerequest}
	jsoned, _ := json.Marshal(all)
	outgoing, _ :=SendRequest(string(jsoned))
	return outgoing
}


func DeleteCustomerProfile(profileID string){
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	profile := deleteCustomerProfileRequest{authToken, profileID}
	input := deleteCustomerProfile{profile}
	jsoned, _ := json.Marshal(input)
	SendRequest(string(jsoned))
}


func CreateCustomerBillingProfile(profileID string, creditCard CreditCard, address Address) (string, map[string]string) {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	paymentProfile := PaymentBillingProfile{Address: address, Payment: Payment{CreditCard:creditCard}}
	request := CreateCustomerBillingProfileRequest{authToken, profileID, paymentProfile, "testMode"}
	newprofile := NewCustomerBillingProfile{request}
	jsoned, _ := json.Marshal(newprofile)
	outgoing, _ :=SendRequest(string(jsoned))
	var new_paymentID string
	var errors map[string]string
	if outgoing["customerPaymentProfileId"]==nil {
		new_paymentID = "0"
		errors = map[string]string{"message": "User cannot be created"}
	} else {
		new_paymentID = outgoing["customerPaymentProfileId"].(string)
		errors = nil
	}
	return new_paymentID, errors
}



func GetCustomerPaymentProfile(profileID string, paymentID string) map[string]interface{} {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	profile := CustomerPaymentProfileRequest{authToken, profileID, paymentID}
	input := getCustomerPaymentProfileRequest{profile}
	jsoned, _ := json.Marshal(input)
	outgoing, _ := SendRequest(string(jsoned))
	return outgoing
}


func UpdateCustomerPaymentProfile(profileID string, paymentID string, address Address, creditCard CreditCard) map[string]interface{} {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	new_billing := UpdatePaymentBillingProfile{Address: address, Payment: Payment{CreditCard:creditCard}, CustomerPaymentProfileId: paymentID}
	profile := updateCustomerPaymentProfileRequest{authToken, profileID, new_billing, "testMode"}
	input := changeCustomerPaymentProfileRequest{profile}
	jsoned, _ := json.Marshal(input)
	outgoing, _ := SendRequest(string(jsoned))
	return outgoing
}


func DeleteCustomerPaymentProfile(profileID string, paymentID string) {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	profile := deleteCustomerPaymentProfile{authToken, profileID, paymentID}
	input := deleteCustomerPaymentProfileRequest{profile}
	jsoned, _ := json.Marshal(input)
	SendRequest(string(jsoned))
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



func CreateTransaction(profileID string, paymentID string, item LineItem, amount string) (map[string]interface{}, map[string]string) {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	items := LineItems{LineItem: item}
	subProfile := SubProfile{CustomerPaymentProfileId: paymentID}
	transProfile := TranProfile{CustomerProfileId: profileID, SubProfile: subProfile}
	transaction := TransactionRequest{TransactionType: "authCaptureTransaction", Amount: amount, TranProfile: transProfile, LineItems: items}
	tranxrequest := CreateTransactionRequest{MerchantAuthentication: authToken, RefID: "none33", TransactionRequest: transaction}
	doTranx := DoCreateTransaction{tranxrequest}
	jsoned, _ := json.Marshal(doTranx)
	outgoing, _ := SendRequest(string(jsoned))
	return outgoing, map[string]string{"yoyo":"yoyook"}
}


func TestConnection() bool {

	authToken := AuthenticateTestRequest{MerchantAuthentication{Name: apiName, TransactionKey: apiKey}}
	authnettest := AuthorizeNetTest{AuthenticateTestRequest:authToken}
	jsoned, _ := json.Marshal(authnettest)
	outgoing, _ := SendRequest(string(jsoned))
	outinner, _ := outgoing["messages"].(map[string]interface{})
	status := outinner["resultCode"]
	if status=="Ok" {
		return true
	}
	return false
}


func CreateShippingAddress(profileID string, address Address) string {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	customerShipping := CustomerShippingAddress{authToken,profileID,address}
	customerShippingRequest := CustomerShippingAddressRequest{customerShipping}
	jsoned, _ := json.Marshal(customerShippingRequest)
	outgoing, _ := SendRequest(string(jsoned))
	var new_address_id string
	if outgoing["customerAddressId"]==nil {
		new_address_id = "0"
	} else {
		new_address_id = outgoing["customerAddressId"].(string)
	}
	return new_address_id
}

func GetShippingAddress(profileID string, shippingID string) map[string]interface{} {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	customerShipping := GetCustomerShippingAddress{authToken,profileID,shippingID}
	customerShippingRequest := GetCustomerShippingAddressRequest{customerShipping}
	jsoned, _ := json.Marshal(customerShippingRequest)
	outgoing, _ := SendRequest(string(jsoned))
	return outgoing
}

func DeleteShippingAddress(profileID string, shippingID string) map[string]interface{} {
	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
	customerShipping := GetCustomerShippingAddress{authToken,profileID,shippingID}
	customerShippingRequest := DeleteCustomerShippingAddressRequest{customerShipping}
	jsoned, _ := json.Marshal(customerShippingRequest)
	outgoing, _ := SendRequest(string(jsoned))
	return outgoing
}


func RefundTransactions(){

}

func VoidTransaction(){

}

func CreateSubscription(){

}

func DeleteSubscription(){

}

func UpdateSubscription(){

}

func GetSubscriptions(){

}

