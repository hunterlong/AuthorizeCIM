package AuthorizeCIM

import (
	"testing"
	"time"
	"math/rand"
	"os"
)


var testProfileID string
var testPaymentID string
var testShippingID string

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
	apiName = os.Getenv("apiName")
	apiKey = os.Getenv("apiKey")
	SetAPIInfo(apiName,apiKey)
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
	testProfileID = newuser
	t.Log("User was created "+testProfileID+"\n")
}


func TestGetCustomerProfile(t *testing.T){
	profile := GetCustomerProfile(testProfileID)
	if profile==nil{
		t.Fail()
	}
	t.Log("Fetched single Customer Profile \n")
}


func TestCreatePaymentProfile(t *testing.T){
	address := Address{FirstName: "Test", LastName: "User", Address: "1234 Road St", City: "City Name", State:" California", Zip: "93063", Country: "USA", PhoneNumber: "5555555555"}
	creditCard := CreditCard{CardNumber: "4111111111111111", ExpirationDate: "2020-12"}
	newPaymentID, err := CreateCustomerBillingProfile(testProfileID, creditCard, address)
	if err!=nil{
		t.Fail()
	}
	testPaymentID = newPaymentID
	t.Log("User Payment Profile created "+testPaymentID+"\n")
}



func TestGetCustomerPaymentProfile(t *testing.T){
	userPaymentProfile := GetCustomerPaymentProfile(testProfileID, testPaymentID)
	if userPaymentProfile==nil {
		t.Fail()
	}
	t.Log("Fetched the Users Payment Profile \n")
}



func TestCreateShippingAddress(t *testing.T){
	address := Address{FirstName: "Test", LastName: "User", Address: "18273 Different St", City: "Los Angeles", State:" California", Zip: "93065", Country: "USA", PhoneNumber: "5555555555"}
	userNewShipping := CreateShippingAddress(testProfileID, address)
	if userNewShipping=="0" {
		t.Fail()
	}
	testShippingID = userNewShipping
	t.Log("Created New Shipping Profile for user "+userNewShipping+"\n")
}


func TestGetShippingAddress(t *testing.T){
	userShipping := GetShippingAddress(testProfileID, testShippingID)
	if userShipping==nil {
		t.Fail()
	}
	t.Log("Fetched User Shipping Address "+testShippingID+"\n")
}


func TestDeleteShippingAddress(t *testing.T){
	userShipping := DeleteShippingAddress(testProfileID, testShippingID)
	if userShipping==nil {
		t.Fail()
	}
	t.Log("Deleted User Shipping Address "+testShippingID+"\n")
}



func TestUpdateCustomerPaymentProfile(t *testing.T){
	address := Address{FirstName: "Test", LastName: "User", Address: "58585 Changed St", City: "Bulbasaur", State:" California", Zip: "93065", Country: "USA", PhoneNumber: "5555555555"}
	creditCard := CreditCard{CardNumber: "4111111111111111", ExpirationDate: "2020-12"}
	updatedPaymentProfile := UpdateCustomerPaymentProfile(testProfileID, testPaymentID, address, creditCard)
	if updatedPaymentProfile==nil {
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
	tranx, _ := CreateTransaction(testProfileID, testPaymentID, item, amount)
	var tranxID string
	fixtransx, _ := tranx["transactionResponse"].(map[string]interface{})
	if tranx["transactionResponse"]==nil {
		tranxID = "0"
		t.Fail()
		t.Log("Transaction has failed! "+tranxID+"\n")
	} else {
		tranxID = fixtransx["transId"].(string)
		t.Log("Submitted and Received Transaction: "+tranxID+"\n")
	}
}


func TestDeleteCustomerPaymentProfile(t *testing.T) {
	DeleteCustomerPaymentProfile(testProfileID, testPaymentID)
	t.Log("User Payment Profile was deleted: "+testPaymentID+"\n")
}

func TestDeleteCustomerProfile(t *testing.T){
	DeleteCustomerProfile(testProfileID)
	t.Log("Customer Profile was deleted: "+testProfileID+"\n")
}