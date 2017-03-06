package AuthorizeCIM

//
//import "encoding/json"
//
//func CreateShippingAddress(profileID string, address Address) (string, bool) {
//	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
//	customerShipping := CustomerShippingAddress{authToken, profileID, address}
//	customerShippingRequest := CustomerShippingAddressRequest{customerShipping}
//	jsoned, _ := json.Marshal(customerShippingRequest)
//	outgoing, _ := SendRequest(string(jsoned))
//	success, _ := FindResultCode(outgoing)
//	var new_address_id string
//	if !success {
//		new_address_id = "0"
//	} else {
//		new_address_id = outgoing["customerAddressId"].(string)
//	}
//	return new_address_id, success
//}
//
//func GetShippingAddress(profileID string, shippingID string) (map[string]interface{}, bool) {
//	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
//	customerShipping := GetCustomerShippingAddress{authToken, profileID, shippingID}
//	customerShippingRequest := GetCustomerShippingAddressRequest{customerShipping}
//	jsoned, _ := json.Marshal(customerShippingRequest)
//	outgoing, _ := SendRequest(string(jsoned))
//	success, _ := FindResultCode(outgoing)
//	return outgoing["address"].(map[string]interface{}), success
//}
//
//func DeleteShippingAddress(profileID string, shippingID string) bool {
//	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
//	customerShipping := GetCustomerShippingAddress{authToken, profileID, shippingID}
//	customerShippingRequest := DeleteCustomerShippingAddressRequest{customerShipping}
//	jsoned, _ := json.Marshal(customerShippingRequest)
//	outgoing, _ := SendRequest(string(jsoned))
//	status, _ := FindResultCode(outgoing)
//	return status
//}
//
//func CreateCustomerProfile(userInfo AuthUser) (string, map[string]interface{}, bool) {
//	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
//	profile := Profile{MerchantCustomerID: userInfo.Uuid, Description: userInfo.Description, Email: userInfo.Email}
//	request := CreateCustomerProfileRequest{authToken, profile}
//	newprofile := NewCustomerProfile{request}
//	jsoned, _ := json.Marshal(newprofile)
//	outgoing, _ := SendRequest(string(jsoned))
//	success, _ := FindResultCode(outgoing)
//	response := outgoing
//	var new_uuid string
//	if success {
//		new_uuid = outgoing["customerProfileId"].(string)
//		CurrentUser.ProfileID = new_uuid
//	} else {
//		new_uuid = "0"
//	}
//	return new_uuid, response, success
//}
//
//func GetCustomerProfile(profileID string) (map[string]interface{}, bool) {
//	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
//	profile := getCustomerProfileRequest{authToken, profileID}
//	input := CustomerProfile{profile}
//	jsoned, _ := json.Marshal(input)
//	outgoing, _ := SendRequest(string(jsoned))
//	success, _ := FindResultCode(outgoing)
//	if outgoing["profile"] == nil {
//		return nil, success
//	} else {
//		userProfile := outgoing["profile"].(map[string]interface{})
//		return userProfile, success
//	}
//}
//
//func GetAllProfiles() []interface{} {
//	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
//	profilerequest := getCustomerProfileIdsRequest{authToken}
//	all := AllCustomerProfileIds{profilerequest}
//	jsoned, _ := json.Marshal(all)
//	outgoing, _ := SendRequest(string(jsoned))
//	return outgoing["ids"].([]interface{})
//}
//
//func DeleteCustomerProfile(profileID string) bool {
//	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
//	profile := deleteCustomerProfileRequest{authToken, profileID}
//	input := deleteCustomerProfile{profile}
//	jsoned, _ := json.Marshal(input)
//	outgoing, _ := SendRequest(string(jsoned))
//	status, _ := FindResultCode(outgoing)
//	return status
//}
//
////func CreateCustomerBillingProfile(profileID string, creditCard CreditCard, address Address) (string, map[string]interface{}, bool) {
////	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
////	paymentProfile := PaymentBillingProfile{Address: address, Payment: Payment{CreditCard: creditCard}}
////	request := CreateCustomerBillingProfileRequest{authToken, profileID, paymentProfile, testMode}
////	newprofile := NewCustomerBillingProfile{request}
////	jsoned, _ := json.Marshal(newprofile)
////	outgoing, _ := SendRequest(string(jsoned))
////	success, _ := FindResultCode(outgoing)
////	response := outgoing
////	var new_paymentID string
////	if success {
////		new_paymentID = outgoing["customerPaymentProfileId"].(string)
////	} else {
////		new_paymentID = "0"
////	}
////	return new_paymentID, response, success
////}
//
//func GetCustomerPaymentProfile(profileID string, paymentID string) (map[string]interface{}, bool) {
//	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
//	profile := CustomerPaymentProfileRequest{authToken, profileID, paymentID}
//	input := getCustomerPaymentProfileRequest{profile}
//	jsoned, _ := json.Marshal(input)
//	outgoing, _ := SendRequest(string(jsoned))
//	success, _ := FindResultCode(outgoing)
//	//fmt.Println(outgoing["paymentProfile"])
//	if success {
//		return outgoing["paymentProfile"].(map[string]interface{}), success
//	} else {
//		//fmt.Println(errMsg)
//		return map[string]interface{}{}, false
//	}
//}
//
////func UpdateCustomerPaymentProfile(profileID string, paymentID string, creditCard CreditCard, address Address) bool {
////	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
////	new_billing := UpdatePaymentBillingProfile{Address: address, Payment: Payment{CreditCard: creditCard}, CustomerPaymentProfileId: paymentID}
////	profile := updateCustomerPaymentProfileRequest{authToken, profileID, new_billing, testMode}
////	input := changeCustomerPaymentProfileRequest{profile}
////	jsoned, _ := json.Marshal(input)
////	outgoing, _ := SendRequest(string(jsoned))
////	status, _ := FindResultCode(outgoing)
////	return status
////}
//
//func DeleteCustomerPaymentProfile(profileID string, paymentID string) bool {
//	authToken := MerchantAuthentication{Name: apiName, TransactionKey: apiKey}
//	profile := deleteCustomerPaymentProfile{authToken, profileID, paymentID}
//	input := deleteCustomerPaymentProfileRequest{profile}
//	jsoned, _ := json.Marshal(input)
//	outgoing, _ := SendRequest(string(jsoned))
//	status, _ := FindResultCode(outgoing)
//	return status
//}
