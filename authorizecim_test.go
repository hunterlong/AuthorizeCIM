package AuthorizeCIM

import (
	"testing"
	"time"
	"math/rand"
	"os"
)


var test_profile_id string
var test_payment_id string
var test_shipping_id string

func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}


func TestSetAPIInfo(t *testing.T) {
	api_name = os.Getenv("api_name")
	api_key = os.Getenv("api_key")
	SetAPIInfo(api_name,api_key)
	t.Log("Authorize.net API Keys have been set! \n")
}

func TestAPIAccess(t *testing.T) {
	if TestConnection() {
		t.Log("API Connection was successful \n")
	} else {
		t.Log("API Connection has incorrect API Keys! \n")
		t.Fail()
	}
}

func TestUserCreation(t *testing.T) {
	customer_info := AuthUser{"54",RandomString(9)+"@random.com","Test Account"}
	newuser, err := CreateCustomerProfile(customer_info)
	if err!=nil{
		t.Fail()
	}
	test_profile_id = newuser
	t.Log("User was created "+test_profile_id+"\n")
}


func TestGetCustomerProfile(t *testing.T){
	profile := GetCustomerProfile(test_profile_id)
	if profile==nil{
		t.Fail()
	}
	t.Log("Fetched single Customer Profile \n")
}


func TestCreatePaymentProfile(t *testing.T){
	address := Address{FirstName: "Test", LastName: "User", Address: "1234 Road St", City: "City Name", State:" California", Zip: "93063", Country: "USA", PhoneNumber: "5555555555"}
	credit_card := CreditCard{CardNumber: "4111111111111111", ExpirationDate: "2020-12"}
	new_payment_id, err := CreateCustomerBillingProfile(test_profile_id, credit_card, address)
	if err!=nil{
		t.Fail()
	}
	test_payment_id = new_payment_id
	t.Log("User Payment Profile created "+test_payment_id+"\n")
}



func TestGetCustomerPaymentProfile(t *testing.T){
	user_payment_profile := GetCustomerPaymentProfile(test_profile_id, test_payment_id)
	if user_payment_profile==nil {
		t.Fail()
	}
	t.Log("Fetched the Users Payment Profile \n")
}



func TestCreateShippingAddress(t *testing.T){
	address := Address{FirstName: "Test", LastName: "User", Address: "18273 Different St", City: "Los Angeles", State:" California", Zip: "93065", Country: "USA", PhoneNumber: "5555555555"}
	user_new_shipping := CreateShippingAddress(test_profile_id, address)
	if user_new_shipping=="0" {
		t.Fail()
	}
	test_shipping_id = user_new_shipping
	t.Log("Created New Shipping Profile for user "+user_new_shipping+"\n")
}


func TestGetShippingAddress(t *testing.T){
	user_shipping := GetShippingAddress(test_profile_id, test_shipping_id)
	if user_shipping==nil {
		t.Fail()
	}
	t.Log("Fetched User Shipping Address "+test_shipping_id+"\n")
}


func TestDeleteShippingAddress(t *testing.T){
	user_shipping := DeleteShippingAddress(test_profile_id, test_shipping_id)
	if user_shipping==nil {
		t.Fail()
	}
	t.Log("Deleted User Shipping Address "+test_shipping_id+"\n")
}



func TestUpdateCustomerPaymentProfile(t *testing.T){
	address := Address{FirstName: "Test", LastName: "User", Address: "58585 Changed St", City: "Bulbasaur", State:" California", Zip: "93065", Country: "USA", PhoneNumber: "5555555555"}
	credit_card := CreditCard{CardNumber: "4111111111111111", ExpirationDate: "2020-12"}
	updated_payment_profile := UpdateCustomerPaymentProfile(test_profile_id, test_payment_id, address, credit_card)
	if updated_payment_profile==nil {
		t.Fail()
	}
	t.Log("Updated the Users Payment Profile with new information \n")
}




func TestGetAllProfiles(t *testing.T){
	profiles := GetAllProfiles()
	if profiles==nil{
		t.Fail()
	}
	t.Log("Fetched ALL Customer Profiles \n")
}




func TestProfileTransaction(t *testing.T) {
	item := LineItem{ItemID: "S0897", Name: "New Product", Description: "brand new", Quantity: "1", UnitPrice: "5.50"}
	amount := "14.43"
	tranx, _ := CreateTransaction(test_profile_id, test_payment_id, item, amount)
	var tranx_id string
	fixtransx, _ := tranx["transactionResponse"].(map[string]interface{})
	if tranx["transactionResponse"]==nil {
		tranx_id = "0"
		t.Fail()
		t.Log("Transaction has failed! "+tranx_id+"\n")
	} else {
		tranx_id = fixtransx["transId"].(string)
		t.Log("Submitted and Received Transaction: "+tranx_id+"\n")
	}
}


func TestDeleteCustomerPaymentProfile(t *testing.T) {
	DeleteCustomerPaymentProfile(test_profile_id, test_payment_id)
	t.Log("User Payment Profile was deleted: "+test_payment_id+"\n")
}

func TestDeleteCustomerProfile(t *testing.T){
	DeleteCustomerProfile(test_profile_id)
	t.Log("Customer Profile was deleted: "+test_profile_id+"\n")
}