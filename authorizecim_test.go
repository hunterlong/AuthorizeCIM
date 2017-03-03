package AuthorizeCIM

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

var testProfileID string
var testPaymentID string
var testShippingID string
var testTransactionID string
var randomUserEmail string
var newSubscriptionId string

func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func RandomDollar(min int, max int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	amount := float64(rand.Intn(max-min) + min)
	f := strconv.FormatFloat(amount, 'f', 2, 64)
	return f
}

func TestSetAPIInfo(t *testing.T) {
	apiName = os.Getenv("apiName")
	apiKey = os.Getenv("apiKey")
	SetAPIInfo(apiName, apiKey, "test")
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
	randomUserEmail = RandomString(9) + "@random.com"
	customer_info := AuthUser{"19", randomUserEmail, "Test Account"}
	newuser, _, success := CreateCustomerProfile(customer_info)
	if !success {
		t.Fail()
	}
	testProfileID = newuser
	t.Log("User was created " + testProfileID + "\n")
}

func TestUserModelCreation(t *testing.T) {
	CurrentUser = MakeUser(randomUserEmail)
	subscriptions := CurrentUser.Subscriptions
	profileid := CurrentUser.ProfileID
	t.Log(subscriptions)
	t.Log(profileid)
}

func TestCreatePaymentProfile(t *testing.T) {
	address := Address{FirstName: "Test", LastName: "User", Address: "1234 Road St", City: "City Name", State: " California", Zip: "93063", Country: "USA", PhoneNumber: "5555555555"}
	creditCard := CreditCard{CardNumber: "4111111111111111", ExpirationDate: "2020-12"}
	newPaymentID, _, success := CreateCustomerBillingProfile(testProfileID, creditCard, address)
	if !success {
		t.Fail()
	}
	testPaymentID = newPaymentID
	t.Log("User Payment Profile created " + testPaymentID + "\n")
}

func TestPaymentDelay(t *testing.T) {
	time.Sleep(15 * time.Second)
	t.Log("Done sleeping \n")
}

func TestGetCustomerPaymentProfile(t *testing.T) {
	userPaymentProfile, success := GetCustomerPaymentProfile(testProfileID, testPaymentID)
	if !success {
		t.Fail()
	}
	t.Log("Fetched the Users Payment Profile \n")
	t.Log(userPaymentProfile)
	t.Log("\n")
}

func TestCreateShippingAddress(t *testing.T) {
	address := Address{FirstName: "Test", LastName: "User", Address: "18273 Different St", City: "Los Angeles", State: " California", Zip: "93065", Country: "USA", PhoneNumber: "5555555555"}
	userNewShipping, success := CreateShippingAddress(testProfileID, address)
	if !success {
		t.Fail()
	}
	testShippingID = userNewShipping
	t.Log("Created New Shipping Profile for user " + userNewShipping + "\n")
}

func TestCreateAnotherShippingAddress(t *testing.T) {
	address := Address{FirstName: "Test", LastName: "User", Address: "18273 MOrse St", City: "Los Nowhere", State: " California", Zip: "87048", Country: "USA", PhoneNumber: "5555555555"}
	userNewShipping, success := CreateShippingAddress(testProfileID, address)
	if !success {
		t.Fail()
	}
	testShippingID = userNewShipping
	t.Log("Created Another Shipping Profile for user " + userNewShipping + "\n")
}

func TestGetShippingAddress(t *testing.T) {
	userShipping, success := GetShippingAddress(testProfileID, testShippingID)
	if !success {
		t.Fail()
	}
	t.Log("Fetched User Shipping Address " + testShippingID + "\n")
	t.Log(userShipping)
	t.Log("\n")
}

func TestGetCustomerProfile(t *testing.T) {
	profile, success := GetCustomerProfile(testProfileID)
	if !success {
		t.Fail()
	}
	t.Log("Fetched single Customer Profile \n")

	t.Log(profile)
	t.Log("\n")
	t.Log("Sleeping for 60 seconds to make sure Auth.net can keep up \n")
}

func TestDelay(t *testing.T) {
	time.Sleep(30 * time.Second)
	t.Log("Done sleeping \n")
}

func TestGetAllProfiles(t *testing.T) {
	profiles := GetAllProfiles()
	if profiles == nil {
		t.Fail()
	}
	t.Log("Fetched ALL Customer Profiles IDs \n")
}

func TestProfileTransaction(t *testing.T) {

	amount := RandomDollar(10, 90)
	item := LineItem{ItemID: "S0592", Name: "New Product", Description: "brand new", Quantity: "1", UnitPrice: amount}
	transResponse, approved, success := CreateTransaction(testProfileID, testPaymentID, item, amount)
	var tranxID string

	if success {
		tranxID = transResponse["transId"].(string)
		testTransactionID = tranxID
		if approved {
			t.Log("Transaction was approved! " + tranxID + "\n")
		} else {
			t.Fail()
			t.Log("Transaction was denied! " + tranxID + "\n")
		}
	} else {
		t.Log("Transaction has failed! It was a duplication transaction or card was rejected. \n")
	}
	t.Log(transResponse)
	t.Log("\n")
}

//func TestAuthorizeCard(t *testing.T) {
//	creditCard := CreditCardCVV{CardNumber: "4012888818888", ExpirationDate: "10/20", CardCode: "433"}
//	approved := AuthorizeCard(creditCard, "1.00")
//	//if approved {
//	//	t.Fail()
//	//}
//}

//func TestRejectedAuthorizeCard(t *testing.T) {
//	creditCard := CreditCardCVV{CardNumber: "401234348888", ExpirationDate: "10/20", CardCode: "433"}
//	approved := AuthorizeCard(creditCard, "1.00")
//	//if !approved {
//	//	t.Fail()
//	//}
//}

func TestProfileTransactionApproved(t *testing.T) {

	// make a new billing profile with a credit card that will be declined
	address := Address{FirstName: "Test", LastName: "User", Address: "1234 Road St", City: "City Name", State: " California", Zip: "93063", Country: "USA", PhoneNumber: "5555555555"}
	creditCard := CreditCard{CardNumber: "4007000000027", ExpirationDate: "2020-12"}
	newPaymentID, _, success := CreateCustomerBillingProfile(testProfileID, creditCard, address)
	if !success {
		t.Fail()
	}
	newTestPaymentID := newPaymentID
	t.Log("New User Payment Profile created " + testPaymentID + "\n")

	// Delay for Authorize.net, waiting for Billing ID
	time.Sleep(10 * time.Second)

	amount := RandomDollar(10, 90)
	item := LineItem{ItemID: "S0595", Name: "New Product", Description: "brand new", Quantity: "1", UnitPrice: amount}
	transResponse, approved, success := CreateTransaction(testProfileID, newTestPaymentID, item, amount)
	var tranxID string

	if success {
		tranxID = transResponse["transId"].(string)
		testTransactionID = tranxID
		if approved {
			t.Log("Transaction was approved! " + tranxID + "\n")
		} else {
			t.Fail()
			t.Log("Transaction was denied! " + tranxID + "\n")
		}
	} else {
		t.Fail()
		t.Log("Transaction has failed! It was a duplication transaction or card was rejected. \n")
	}
	t.Log(transResponse)
	t.Log("\n")
}

func TestGetTransactionDetails(t *testing.T) {
	details := GetTransactionDetails(testTransactionID)
	if details != nil {
		if details["transId"] == testTransactionID {
			t.Log("Transaction ID " + testTransactionID + " was fetched! \n")
		} else {
			t.Fail()
			t.Log("Transaction was not processed! Could be a duplicate transaction. \n")
		}
	}
	t.Log(details)
}

func TestCreateSubscription(t *testing.T) {

	startTime := time.Now().Format("2006-01-02")
	//startTime := "2016-06-02"
	totalRuns := "9999" //means forever
	trialRuns := "0"
	amount := RandomDollar(10, 90)
	userFullProfile := FullProfile{CustomerProfileID: testProfileID, CustomerAddressID: testShippingID, CustomerPaymentProfileID: testPaymentID}
	paymentSchedule := PaymentSchedule{Interval: Interval{"1", "months"}, StartDate: startTime, TotalOccurrences: totalRuns, TrialOccurrences: trialRuns}
	subscriptionInput := Subscription{"New Subscription", paymentSchedule, amount, "0.00", userFullProfile}

	newSubscription, success := CreateSubscription(subscriptionInput)
	newSubscriptionId = newSubscription
	if success {
		t.Log("User created a new Subscription id: " + newSubscription + "\n")
	} else {
		t.Fail()
		t.Log("created the subscription failed, the user might not be fully inputed yet. \n")
	}
}

func TestCancelSubscription(t *testing.T) {

	thisSubscriptionId := newSubscriptionId
	success := DeleteSubscription(thisSubscriptionId)

	if success {
		fmt.Print("Canceled Subscription ID: ", thisSubscriptionId)
	} else {
		fmt.Println("Failed to cancel subscription ID: ", thisSubscriptionId)
		t.Fail()
	}

}

func TestRefundTransaction(t *testing.T) {

	transId := "10102012"
	amount := "10.50"
	lastFour := "4040"
	expiration := "10/20"

	response, approved, status := RefundTransaction(transId, amount, lastFour, expiration)

	if status {

		if approved {
			fmt.Println("Transaction Refund was APPROVED and processed")
		} else {
			fmt.Println("Transaction Refund was decliened")
		}

	} else {
		fmt.Println("Refund failed to process")
	}
	fmt.Println(response)

}

func TestUpdateCustomerPaymentProfile(t *testing.T) {
	address := Address{FirstName: "Test", LastName: "User", Address: "58585 Changed St", City: "Bulbasaur", State: " California", Zip: "93065", Country: "USA", PhoneNumber: "5555555555"}
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

func TestDeleteShippingAddress(t *testing.T) {
	userShipping := DeleteShippingAddress(testProfileID, testShippingID)
	if userShipping {
		t.Log("Deleted User Shipping Address " + testShippingID + "\n")
	} else {
		t.Log("Issue with deleteing shippinn address: " + testShippingID + "\n")
	}
}

func TestDeleteCustomerProfile(t *testing.T) {
	response := DeleteCustomerProfile(testProfileID)
	if response {
		t.Log("Customer Profile was deleted: " + testProfileID + "\n")
	}
}

func TestVoidTransaction(t *testing.T) {

	success := VoidTransaction(testTransactionID)

	if success {
		fmt.Println("Transaction was successfully voided")
	} else {
		fmt.Println("Transaction FAILED to void")
	}
}
