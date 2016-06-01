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
var testTransactionID string

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
	customer_info := AuthUser{"73",RandomString(9)+"@random.com","Test Account"}
	newuser, success := CreateCustomerProfile(customer_info)
	if !success {
		t.Fail()
	}
	testProfileID = newuser
	t.Log("User was created "+testProfileID+"\n")
}


func TestCreatePaymentProfile(t *testing.T){
	address := Address{FirstName: "Test", LastName: "User", Address: "1234 Road St", City: "City Name", State:" California", Zip: "93063", Country: "USA", PhoneNumber: "5555555555"}
	creditCard := CreditCard{CardNumber: "4111111111111111", ExpirationDate: "2020-12"}
	newPaymentID, success := CreateCustomerBillingProfile(testProfileID, creditCard, address)
	if !success {
		t.Fail()
	}
	testPaymentID = newPaymentID
	t.Log("User Payment Profile created "+testPaymentID+"\n")}



func TestGetCustomerPaymentProfile(t *testing.T){
	userPaymentProfile, success := GetCustomerPaymentProfile(testProfileID, testPaymentID)
	if !success {
		t.Fail()
	}
	t.Log("Fetched the Users Payment Profile \n")
	t.Log(userPaymentProfile)
	t.Log("\n")
}



func TestCreateShippingAddress(t *testing.T){
	address := Address{FirstName: "Test", LastName: "User", Address: "18273 Different St", City: "Los Angeles", State:" California", Zip: "93065", Country: "USA", PhoneNumber: "5555555555"}
	userNewShipping, success := CreateShippingAddress(testProfileID, address)
	if !success {
		t.Fail()
	}
	testShippingID = userNewShipping
	t.Log("Created New Shipping Profile for user "+userNewShipping+"\n")
}


func TestGetShippingAddress(t *testing.T){
	userShipping, success := GetShippingAddress(testProfileID, testShippingID)
	if !success {
		t.Fail()
	}
	t.Log("Fetched User Shipping Address "+testShippingID+"\n")
	t.Log(userShipping)
	t.Log("\n")
}


func TestGetAllProfiles(t *testing.T){
	profiles := GetAllProfiles()
	if profiles==nil{
		t.Fail()
	}
	t.Log("Fetched ALL Customer Profiles IDs \n")
}




func TestProfileTransaction(t *testing.T) {
	item := LineItem{ItemID: "S0891", Name: "New Product", Description: "brand new", Quantity: "1", UnitPrice: "8.50"}
	amount := "18.28"
	transResponse, approved, success := CreateTransaction(testProfileID, testPaymentID, item, amount)
	var tranxID string

	if success {
		tranxID = transResponse["transId"].(string)
		testTransactionID = tranxID
		if approved {
			t.Log("Transaction was approved! "+tranxID+"\n")
		} else {
			t.Log("Transaction was denied! "+tranxID+"\n")
		}
	} else {
		t.Log("Transaction has failed! It was a duplication transaction or card was rejected. \n")
	}
	t.Log(transResponse)
	t.Log("\n")
}


func TestGetTransactionDetails(t *testing.T) {
	details := GetTransactionDetails(testTransactionID)
	if details != nil {
	if details["transId"] == testTransactionID {
		t.Log("Transaction ID "+testTransactionID+" was fetched! \n")
	} else {
		t.Log("Transaction was not processed! Could be a duplicate transaction. \n")
	}
	}
}


func TestGetCustomerProfile(t *testing.T){
	profile, success := GetCustomerProfile(testProfileID)
	if !success {
		t.Fail()
	}
	t.Log("Fetched single Customer Profile \n")
	t.Log(profile)
	t.Log("\n")
	t.Log("Sleeping for 60 seconds to make sure Auth.net can keep up \n")
}



func TestDelay(t *testing.T){
	time.Sleep(60 * time.Second)
	t.Log("Done sleeping \n")
}



func TestCreateSubscription(t *testing.T){

	startTime := time.Now().Format("2006-01-02")
	//startTime := "2016-06-02"
	totalRuns := "9999" //means forever
	trialRuns := "0"
	userFullProfile := FullProfile{CustomerProfileID: testProfileID,CustomerAddressID: testShippingID, CustomerPaymentProfileID: testPaymentID}
	paymentSchedule := PaymentSchedule{Interval: Interval{"1","months"}, StartDate:startTime, TotalOccurrences:totalRuns, TrialOccurrences:trialRuns}
	subscriptionInput := Subscription{"Advanced Subscription",paymentSchedule,"7.98","0.00",userFullProfile}

	newSubscription, success := CreateSubscription(subscriptionInput)
	if success {
		t.Log("User created a new Subscription id: "+newSubscription+"\n")
	} else {
		t.Log("created the subscription failed, the user might not be fully inputed yet. \n")
	}
}




func TestUpdateCustomerPaymentProfile(t *testing.T){
	address := Address{FirstName: "Test", LastName: "User", Address: "58585 Changed St", City: "Bulbasaur", State:" California", Zip: "93065", Country: "USA", PhoneNumber: "5555555555"}
	creditCard := CreditCard{CardNumber: "4111111111111111", ExpirationDate: "2020-12"}
	updatedPaymentProfile := UpdateCustomerPaymentProfile(testProfileID, testPaymentID, creditCard, address)
	if !updatedPaymentProfile {
		t.Fail()
	}
	t.Log("Updated the Users Payment Profile with new information \n")
}


func TestDeleteCustomerPaymentProfile(t *testing.T) {
	response := DeleteCustomerPaymentProfile(testProfileID, testPaymentID)
	if response {
		t.Log("User Payment Profile was deleted: " + testPaymentID + "\n")
	}
}

func TestDeleteShippingAddress(t *testing.T){
	userShipping := DeleteShippingAddress(testProfileID, testShippingID)
	if userShipping {
		t.Log("Deleted User Shipping Address "+testShippingID+"\n")
	} else {
		t.Log("Issue with deleteing shippinn address: "+testShippingID+"\n")
	}
}


func TestDeleteCustomerProfile(t *testing.T){
	response := DeleteCustomerProfile(testProfileID)
	if response {
		t.Log("Customer Profile was deleted: " + testProfileID + "\n")
	}
}