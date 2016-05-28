package AuthorizeCIM

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

var api_endpoint string = "https://apitest.authorize.net/xml/v1/request.api"
var api_name string
var api_key string


func SetAPIInfo(api_name string, api_key string) {
	api_key = api_key
	api_name = api_name
}

func CreateCustomerProfile(user_info AuthUser) (string, map[string]string) {
	auth_token := MerchantAuthentication{Name: api_name, TransactionKey: api_key}
	profile := Profile{MerchantCustomerID: user_info.Uuid, Description: user_info.Description, Email: user_info.Email}
	request := CreateCustomerProfileRequest{auth_token, profile}
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


func GetCustomerProfile(profile_id string) map[string]interface{} {
	auth_token := MerchantAuthentication{Name: api_name, TransactionKey: api_key}
	profile := getCustomerProfileRequest{auth_token, profile_id}
	input := CustomerProfile{profile}
	jsoned, _ := json.Marshal(input)
	outgoing, _ :=SendRequest(string(jsoned))
	return outgoing
}


func GetAllProfiles() map[string]interface{} {
	auth_token := MerchantAuthentication{Name: api_name, TransactionKey: api_key}
	profilerequest := getCustomerProfileIdsRequest{auth_token}
	all := AllCustomerProfileIds{profilerequest}
	jsoned, _ := json.Marshal(all)
	outgoing, _ :=SendRequest(string(jsoned))
	return outgoing
}


func DeleteCustomerProfile(profile_id string){
	auth_token := MerchantAuthentication{Name: api_name, TransactionKey: api_key}
	profile := deleteCustomerProfileRequest{auth_token, profile_id}
	input := deleteCustomerProfile{profile}
	jsoned, _ := json.Marshal(input)
	SendRequest(string(jsoned))
}


func CreateCustomerBillingProfile(profile_id string, credit_card CreditCard, address Address) (string, map[string]string) {
	auth_token := MerchantAuthentication{Name: api_name, TransactionKey: api_key}
	payment_profile := PaymentBillingProfile{Address: address, Payment: Payment{CreditCard:credit_card}}
	request := CreateCustomerBillingProfileRequest{auth_token, profile_id, payment_profile, "testMode"}
	newprofile := NewCustomerBillingProfile{request}
	jsoned, _ := json.Marshal(newprofile)
	outgoing, _ :=SendRequest(string(jsoned))
	var new_payment_id string
	var errors map[string]string
	if outgoing["customerPaymentProfileId"]==nil {
		new_payment_id = "0"
		errors = map[string]string{"message": "User cannot be created"}
	} else {
		new_payment_id = outgoing["customerPaymentProfileId"].(string)
		errors = nil
	}
	return new_payment_id, errors
}



func GetCustomerPaymentProfile(profile_id string, payment_id string) map[string]interface{} {
	auth_token := MerchantAuthentication{Name: api_name, TransactionKey: api_key}
	profile := CustomerPaymentProfileRequest{auth_token, profile_id, payment_id}
	input := getCustomerPaymentProfileRequest{profile}
	jsoned, _ := json.Marshal(input)
	outgoing, _ := SendRequest(string(jsoned))
	return outgoing
}


func UpdateCustomerPaymentProfile(profile_id string, payment_id string, address Address, credit_card CreditCard) map[string]interface{} {
	auth_token := MerchantAuthentication{Name: api_name, TransactionKey: api_key}
	new_billing := UpdatePaymentBillingProfile{Address: address, Payment: Payment{CreditCard:credit_card}, CustomerPaymentProfileId: payment_id}
	profile := updateCustomerPaymentProfileRequest{auth_token, profile_id, new_billing, "testMode"}
	input := changeCustomerPaymentProfileRequest{profile}
	jsoned, _ := json.Marshal(input)
	outgoing, _ := SendRequest(string(jsoned))
	return outgoing
}


func DeleteCustomerPaymentProfile(profile_id string, payment_id string) {
	auth_token := MerchantAuthentication{Name: api_name, TransactionKey: api_key}
	profile := deleteCustomerPaymentProfile{auth_token, profile_id, payment_id}
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



func CreateTransaction(profile_id string, payment_id string, item LineItem, amount string) (map[string]interface{}, map[string]string) {
	auth_token := MerchantAuthentication{Name: api_name, TransactionKey: api_key}
	items := LineItems{LineItem: item}
	sub_profile := SubProfile{CustomerPaymentProfileId: payment_id}
	trans_profile := TranProfile{CustomerProfileId: profile_id, SubProfile: sub_profile}
	transaction := TransactionRequest{TransactionType: "authCaptureTransaction", Amount: amount, TranProfile: trans_profile, LineItems: items}
	tranxrequest := CreateTransactionRequest{MerchantAuthentication: auth_token, RefID: "none33", TransactionRequest: transaction}
	do_tranx := DoCreateTransaction{tranxrequest}
	jsoned, _ := json.Marshal(do_tranx)
	outgoing, _ := SendRequest(string(jsoned))
	return outgoing, map[string]string{"yoyo":"yoyook"}
}


func TestConnection() bool {

	auth_token := AuthenticateTestRequest{MerchantAuthentication{Name: api_name, TransactionKey: api_key}}
	authnettest := AuthorizeNetTest{AuthenticateTestRequest:auth_token}
	jsoned, _ := json.Marshal(authnettest)
	outgoing, _ := SendRequest(string(jsoned))
	outinner, _ := outgoing["messages"].(map[string]interface{})
	status := outinner["resultCode"]
	if status=="Ok" {
		return true
	}
	return false
}


func CreateShippingAddress(){

}

func DeleteShippingAddress(){

}

func UsersShippingAddresses() {

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

